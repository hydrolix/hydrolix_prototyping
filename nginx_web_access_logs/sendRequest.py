import requests

# very simple request call using python to send csv data - 
#    in production, this would usually be done in high volume with a tool ro via batch or 
#
# prereqs:
#            python
#            request library
# to try it out, just update:
#
#            x-hdx-table with your projectname.tablename 
#            nginx_access_logs with your transform name
#
# To learn more about stream ingest, go here: 
#
#      docs - https://docs.hydrolix.io/docs/streaming-http
#      API -  https://docs.hydrolix.io/reference/stream-event
#
# Security is assumed to be turuned off.  To learn how to turn that on for a table, go here:
#
#      stream authentication - https://docs.hydrolix.io/docs/stream-authentication

hostname = ""; ##enter your DNS within the quotes before running 

url = "https://" + hostname + "/ingest/event"

print("sending to: " + url)

payload = "40.77.167.129 - - [22/Jan/2023:03:56:18 +0330] \"GET /image/45437/productModel/150x150 HTTP/1.1\" 200 3688 \"-\" \"Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)\" \"-\""


headers = {
    "x-hdx-table": "hydrolix_demo_projects.nginx",
    "x-hdx-transform": "nginx_access_logs",
    "content-type": "text/csv",
}

response = requests.post(url, data=payload, headers=headers)

print(response.text)
