

Sample transforms for common log patterns

Each transform ships with

(1) A transform in JSON format which can be uploaded into 


nginx access logs
apache logs
netflow V9 sample

coming soon:
firewall logs 
linux syslog 
window event logs 
JSON IoT Example 


nginx access logs
link_to_source: https://archive.ics.uci.edu/ml/datasets/Online+Retail+II
LIcense CC0: Public Domain

40.77.167.129 - - [22/Jan/2019:03:56:18 +0330] "GET /image/45437/productModel/150x150 HTTP/1.1" 200 3688 "-" "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)" "-"
40.77.167.129 - - [22/Jan/2019:03:56:18 +0330] "GET /image/576/article/100x100 HTTP/1.1" 200 14776 "-" "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)" "-"
66.249.66.194 - - [22/Jan/2019:03:56:18 +0330] "GET /filter/b41,b665,c150%7C%D8%A8%D8%AE%D8%A7%D8%B1%D9%BE%D8%B2,p56 HTTP/1.1" 200 34277 "-" "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)" "-"
40.77.167.129 - - [22/Jan/2019:03:56:18 +0330] "GET /image/57710/productModel/100x100 HTTP/1.1" 200 1695 "-" "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)" "-"



apache logs

https://www.kaggle.com/datasets/vishnu0399/server-logs?select=logfiles.log
CCO license

233.223.117.90 - - [27/Dec/2037:12:00:00 +0530] "DELETE /usr/admin HTTP/1.0" 502 4963 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4380.0 Safari/537.36 Edg/89.0.759.0" 45
162.253.4.179 - - [27/Dec/2037:12:00:00 +0530] "GET /usr/admin/developer HTTP/1.0" 200 5041 "http://www.parker-miller.org/tag/list/list/privacy/" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36" 3885
252.156.232.172 - - [27/Dec/2037:12:00:00 +0530] "POST /usr/register HTTP/1.0" 404 5028 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36 OPR/73.0.3856.329" 3350
182.215.249.159 - - [27/Dec/2037:12:00:00 +0530] "PUT /usr/register HTTP/1.0" 304 4936 "http://www.parker-miller.org/tag/list/list/privacy/" "Mozilla/5.0 (Android 10; Mobile; rv:84.0) Gecko/84.0 Firefox/84.0" 767
160.36.208.51 - - [27/Dec/2037:12:00:00 +0530] "POST /usr HTTP/1.0" 304 4979 "http://www.parker-miller.org/tag/list/list/privacy/" "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6000) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Mobile Safari/537.36 OPR/61.2.3076.56749" 84


netflow V9 sample
https://www.kaggle.com/datasets/ashtcoder/network-data-schema-in-the-netflow-v9-format
Attribution 4.0 International (CC BY 4.0) license

FLOW_ID,PROTOCOL_MAP,L4_SRC_PORT,IPV4_SRC_ADDR,L4_DST_PORT,IPV4_DST_ADDR,FIRST_SWITCHED,FLOW_DURATION_MILLISECONDS,LAST_SWITCHED,PROTOCOL,TCP_FLAGS,TCP_WIN_MAX_IN,TCP_WIN_MAX_OUT,TCP_WIN_MIN_IN,TCP_WIN_MIN_OUT,TCP_WIN_MSS_IN,TCP_WIN_SCALE_IN,TCP_WIN_SCALE_OUT,SRC_TOS,DST_TOS,TOTAL_FLOWS_EXP,MIN_IP_PKT_LEN,MAX_IP_PKT_LEN,TOTAL_PKTS_EXP,TOTAL_BYTES_EXP,IN_BYTES,IN_PKTS,OUT_BYTES,OUT_PKTS,ANALYSIS_TIMESTAMP,ANOMALY,ALERT,ID
372018590,udp,53950,10.114.232.40,53,10.114.226.5,1647760833,1,1647760833,17,0,0,0,0,0,0,0,0,0,0,372018590,0,0,0,0,165,2,275,2,1647760873,0.0,None,0
368604472,tcp,37914,10.114.241.166,38303,10.114.224.218,1647686725,1,1647686725,6,22,1024,0,1024,0,1460,0,0,0,0,368604472,0,0,0,0,44,1,40,1,1647687338,0.0,Port Scanning,1
367513992,tcp,33216,10.114.241.166,18757,10.114.224.116,1647680044,0,1647680044,6,22,1024,0,1024,0,1460,0,0,0,0,367513992,0,0,0,0,44,1,40,1,1647680115,1.0,Port Scanning,2
371997670,udp,48627,10.114.225.205,53,10.114.226.5,1647760172,0,1647760172,17,0,0,0,0,0,0,0,0,0,0,371997670,0,0,0,0,128,2,160,2,1647760207,0.0,None,3
374311119,udp,35939,10.114.225.205,53,10.114.226.5,1647780916,0,1647780916,17,0,0,0,0,0,0,0,0,0,0,374311119,0,0,0,0,172,2,300,2,1647780984,0.0,None,4
338500951,tcp,57175,45.134.144.144,2443,10.114.226.5,1647412843,0,1647412843,6,2,1024,0,1024,0,0,0,0,40,0,338500951,0,0,0,0,40,1,0,0,1647415742,,None,5
338612313,udp,16648,10.114.225.205,53,10.114.226.5,1647414417,0,1647414417,17,0,0,0,0,0,0,0,0,0,0,338612313,0,0,0,0,154,2,418,2,1647421601,,None,6
334640917,udp,63880,10.114.225.216,53,10.114.226.5,1647344697,0,1647344697,17,0,0,0,0,0,0,0,0,0,0,334640917,0,0,0,0,140,2,172,2,1647344753,,None,7


Soon to come:
internet firewall logs  todo
https://www.kaggle.com/datasets/tunguz/internet-firewall-data-set


sample syslogs
https://github.com/logpai/loghub/tree/master/Linux
(freely available)