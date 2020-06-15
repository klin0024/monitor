# Overview
![image](img/overview.JPG?raw=true "Overview") <br />
<br />
## 使用說明

### 1. 使用superuser的權限執行 install
$ sudo /bin/bash install <br />
### 2. 修改monitor.list, 監控目錄清單
$ vi monitor.list <br />
範例 <br />
\#------------------------ <br />
\#ADD MONITORING DIRECTORY <br />
\#------------------------ <br />
/etc/ <br />
\#/home/ <br />
### 3. 執行baseline, 建立基準線
$ baseline <br />
### 4. 執行monitor, 監控程式會產生 checkout.json
$ monitor <br />
<br />
### checkout.json 格式:
[ <br />
&nbsp;&nbsp;{ <br />
&nbsp;&nbsp;&nbsp;&nbsp;"file": "< file path >" (str) <br />
&nbsp;&nbsp;&nbsp;&nbsp;"uid": "< uid >" (str) <br />
&nbsp;&nbsp;&nbsp;&nbsp;"gid": "< gid >" (str) <br />
&nbsp;&nbsp;&nbsp;&nbsp;"modtime": "< unix time(s)>" (str) <br />
&nbsp;&nbsp;&nbsp;&nbsp;"state": "< 0:delete , 1:add , 2:change >" (int) <br />
&nbsp;&nbsp;}, <br />
… <br />
]<br />
### checkout.json 範例:
[ <br />
&nbsp;&nbsp;{ <br />
&nbsp;&nbsp;&nbsp;&nbsp;"file":"/etc/add 1", <br />
&nbsp;&nbsp;&nbsp;&nbsp;"uid":"0", <br />
&nbsp;&nbsp;&nbsp;&nbsp;"gid":"0", <br />
&nbsp;&nbsp;&nbsp;&nbsp;"modtime":"1592185615", <br />
&nbsp;&nbsp;&nbsp;&nbsp;"state":1 <br />
&nbsp;&nbsp;}, <br />
&nbsp;&nbsp;{ <br />
&nbsp;&nbsp;&nbsp;&nbsp;"file":"/etc/remove 1", <br />
&nbsp;&nbsp;&nbsp;&nbsp;"uid":"0", <br />
&nbsp;&nbsp;&nbsp;&nbsp;"gid":"0", <br />
&nbsp;&nbsp;&nbsp;&nbsp;"modtime":"1592179200", <br />
&nbsp;&nbsp;&nbsp;&nbsp;"state":0 <br />
&nbsp;&nbsp;}, <br />
&nbsp;&nbsp;{ <br />
&nbsp;&nbsp;&nbsp;&nbsp;"file":"/etc/change 1", <br />
&nbsp;&nbsp;&nbsp;&nbsp;"uid":"0", <br />
&nbsp;&nbsp;&nbsp;&nbsp;"gid":"0", <br />
&nbsp;&nbsp;&nbsp;&nbsp;"modtime":"1592185671", <br />
&nbsp;&nbsp;&nbsp;&nbsp;"state":2 <br />
&nbsp;&nbsp;} <br />
] <br />