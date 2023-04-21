#########################
####
#### sample query for rows 
####    notice table is referred to by projectname.tablename in the query
####    a best practice is to always limit the window of data being queries in order to economy resource
####    Hydrolix has circuit breakers which can be put in to limit queries across the entire cluster
####
########################
select
    *
from
    hydrolix_demo_projects.nginx
limit
    10 
    
    #query within a time range
    
########################
####
####   now we are going big - with our 1 to 5 records (in reality you will have billions of records and will want to filter your query)
####
########################
select
    *
from
    hydrolix_demo_projects.nginx
WHERE
    primary < now()
    AND primary > now() - INTERVAL 1 HOUR
LIMIT
    5000 

########################
#####
#####    
#####  run an aggregration to count total incoming bytes for an hour
#####
########################
select
    sum(request_size) as sum_of_requests_in_hour
from
    hydrolix_demo_projects.nginx
WHERE
    primary < now()
    AND primary > now() - INTERVAL 1 HOUR
LIMIT
    5000 #search for web logs

########################
#####
#####  check your hydrolix logs for recent activity - there is a great screen to look at system health in the UI, too!
#####   try describing your logs to see the other attributes
#####
########################
select
    timestamp,
    app,
    message,
    catchall
from
    hydro.logs
WHERE
    timestamp < now()
    AND timestamp > now() - INTERVAL 5 MINUTE
order by
    timestamp desc
LIMIT
    1000

########################
#####
##### sample curl to query
##### for more parameters, go here: https://docs.hydrolix.io/reference/query-data-get
####  for more sample curls and request calls, look at the sample_API_calls.http file
#####
########################
curl --request POST \
     --url {{hostname}}/query \
     --header 'accept: text/json' \
     --header 'content-type: text/plain' \
     --header 'authorization: {{apikey}}' \
     --data 'select * from  hydrolix_demo_projects.nginx limit  10 '
####
########################