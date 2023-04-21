

This nginx logs examples shows how to easily:

(1) Build out projects, tables, and transforms in Hydrolix
(2) Model transforms of tables in order to define field types which a Hydrolix table will store
(3) Ingest data using HTTP streaming into Hydrolix with a simple Python program-based POST
(4) Query data within Hydrolix - from Grafana using the Clickhouse client or from the Hydrolix query console, for example

Here is a description of files:

(1) buildtable.sh - sample script for automating building a project, table, and transform with HDXCLI
(2) nginx_transform_description - review of command line how to do the same as (1)
                                review of how to manually build out a transform using the Hydrolix user interface
                                review of design decisions made for the simple NGINX transform used as an example
(3) nginx_access_logs.json - the transform in JSON format 
                           this format will work wtih hdxcli to upload a transform for a table
(4) sendRequest.py - a very simple (hard coded on purpose!) request to ingest data using Python
                   in production, most users stream using solutions like vector.dev, fluentd, kinesis, kafka
(5) sample_queries.sql - sample queries which can be run to queries huge datasets:
                       in Grafana via the Clickhouse client
                       via the Hydrolix UI query screen 
                       via a POST (see item (6) )
(6) sample_API_calls.http - works nicely in VS CODE with HTTP plugin 
                            REST calls which are very useful to run in order to: 
                                 retrieve an authentication token
                                 GET projects, tables, transforms, dictionaries
                                 Create table
                                 Ingest data
                                 Query data

