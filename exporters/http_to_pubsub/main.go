// This example shows how to pull data periodically from the Hydrolix HTTP API
// and publish messages containing that data to Google Pub/Sub. See the block of
// constants below for settings for your environment. A full write-up of the
// methods used here is available at
// https://docs.hydrolix.io/docs/export-pubsub.
//
// Note that this is an approach for low-throughput tables. Situations with
// higher throughput could require non-blocking posts, tracking the size of
// data, parallelization, and other features for more robust data delivery.
package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloud.google.com/go/pubsub"
)

const (
	defaultQueryLoopDuration    = 5 * time.Second
	defaultInitialDelayDuration = 10 * time.Second
	dateFormatString            = "2006-01-02 15:04:05.0000"
)

var (
	hdxQueryAPIURL    string
	hdxAPIToken       string
	hdxTableName      string
	queryLoopDuration time.Duration
	queryInitialDelay time.Duration
	gcpProjectID      string
	gcpTopicName      string
	hdxAPIUser        string
	hdxAPIPassword    string
)

func init() {
	flag.StringVar(&hdxQueryAPIURL, "hdx_url", "", "Hydrolix query URL")
	flag.StringVar(&hdxAPIToken, "hdx_token", "", "Hydrolix API endpoint bearer token")
	flag.StringVar(&hdxTableName, "hdx_table", "", "Hydrolix table to query (in PROJECT.TABLE notation)")
	flag.StringVar(&gcpProjectID, "gcp_project", "", "The Google Cloud Platform project ID to use")
	flag.StringVar(&gcpTopicName, "gcp_topic", "", "The destination topic name (will attempt to create if nonexistent)")
	flag.DurationVar(&queryLoopDuration, "loop_delay", defaultQueryLoopDuration, "The cadence on which to query the HDX table")
	flag.DurationVar(&queryInitialDelay, "initial_delay", defaultInitialDelayDuration, "The duration to wait before querying the table for the first time")
	flag.StringVar(&hdxAPIUser, "hdx_user", "", "The API user")
	flag.StringVar(&hdxAPIPassword, "hdx_password", "", "The API password")
}

func parseArgs() {
	flag.Parse()

	req := func(name, value string) {
		if value == "" {
			log.Fatalf("Parameter --%s must be set.", name)
		}
	}

	req("url", hdxQueryAPIURL)
	req("table", hdxTableName)
	req("gcp_project", gcpProjectID)
	req("gcp_topic", gcpTopicName)

	if hdxAPIToken == "" && (hdxAPIUser == "" || hdxAPIPassword == "") {
		log.Fatalf("Need either --hdx_token or both --hdx_user and --hdx_password set to access API.")
	}

}

// hdxRequest performs the Hydrolix HTTP API request
func hdxRequest(client *http.Client, myQuery string) []byte {
	req, err := http.NewRequest("POST", hdxQueryAPIURL, bytes.NewBuffer([]byte(myQuery)))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set the API token or Username:Password
	if hdxAPIToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", hdxAPIToken))
	} else {
		req.SetBasicAuth(hdxAPIUser, hdxAPIPassword)
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending Hydrolix API request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected Hydrolix API status code: %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)

	log.Printf("\n%s\n", body)
	return body
}

// calculateSQLQuery creates the SQL query needed for the current time frame
func calculateSQLQuery(start, end time.Time, table string) (string, time.Time) {
	startString := start.Format(dateFormatString)
	endString := end.Format(dateFormatString)

	queryString := fmt.Sprintf("SELECT * FROM %s WHERE timestamp >= '%s' AND timestamp < '%s' ORDER BY timestamp DESC FORMAT JSONEachRow", table, startString, endString)
	log.Println(queryString)

	return queryString, end
}

// findOrCreateTopic returns a topic, creating a new one or finding an existing one
func findOrCreateTopic(ctx context.Context, client *pubsub.Client) *pubsub.Topic {

	// Check to see if the topic already exists
	topic := client.Topic(gcpTopicName)
	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatalf("Couldn't check to see if topic %s exists.\n", gcpTopicName)
	}
	if !exists {
		log.Printf("Creating topic %s.\n", gcpTopicName)
		// Creates the new topic.
		topic, err := client.CreateTopic(ctx, gcpTopicName)
		if err != nil {
			log.Printf("Failed to create topic: %v", err)
		}
		log.Printf("Topic %v created.\n", topic)
	} else {
		log.Printf("Topic %s exists.", topic.String())
	}
	return topic
}

func main() {

	parseArgs()

	ctx := context.Background()
	httpClient := &http.Client{}

	// Create the GCP Pub/Sub Client
	client, err := pubsub.NewClient(ctx, gcpProjectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	gcpTopic := findOrCreateTopic(ctx, client)

	osChannel := make(chan os.Signal, 1)
	signal.Notify(osChannel, os.Interrupt, syscall.SIGTERM)

	// Set up a ticker for the duration specified
	ticker := time.NewTicker(queryLoopDuration)
	defer ticker.Stop()

	query := ""

	// Back-date the dates in our query to ensure all data has been received
	// before the query is executed.
	start := time.Now().Add(-queryInitialDelay).UTC()

	// Calculate warning threshold for the loop below so we can log a warning if
	// the loop takes too long.
	warnDelay := (2 * queryLoopDuration) - queryInitialDelay

	for {
		select {
		case <-osChannel:
			log.Println("Exiting.")
			gcpTopic.Stop()
			return
		case t := <-ticker.C:
			// Are we delayed by too much?
			if t.Sub(start) > warnDelay {
				log.Printf("Warning: delayed to %v", warnDelay)
				warnDelay *= 2
			}
			// Create SQL query text and get new start time for next iteration
			query, start = calculateSQLQuery(start, start.Add(queryLoopDuration), hdxTableName)

			// Get data from Hydrolix
			data := hdxRequest(httpClient, query)

			// Publish data to Google Pub/Sub
			result := gcpTopic.Publish(ctx, &pubsub.Message{Data: data})

			// Block for Pub/Sub to finish
			result.Get(ctx)
		}
	}
}
