# Overview <br />
![image](img/overview.JPG?raw=true "Overview") <br />
<br />
使用說明 <br />

1. 使用superuser的權限執行install <br />
$ sudo /bin/bash install <br />
<br />
2. 修改監控目錄清單 monitor.list <br />
$ vi monitor.list <br />
<br />
範例 <br />
#------------------------ <br />
#ADD MONITORING DIRECTORY <br />
#------------------------ <br />
/etc/ <br />
#/home/ <br />
<br />
3. 執行baseline,建立基準線 <br />
$ baseline <br />
<br />
4. 執行monitor,監控程式會產生checkout.json <br />
$ monitor <br />
<br />
checkout.json 格式: <br />
<br />
[ <br />
{ <br />
   "file": "<file path>" (str) <br />
   "uid": "<uid>" (str) <br />
   "gid": "<gid>" (str) <br />
   "modtime": "<unix time(s)>" (str) <br />
   "state": "< 0:delete , 1:add , 2:change >" (int) <br />
}, <br />
… <br />
]<br />
<br />
checkout.json 範例: <br />
<br />
[ <br />
   { <br />
      "file":"/etc/add 1", <br />
      "uid":"0", <br />
      "gid":"0", <br />
      "modtime":"1592185615", <br />
      "state":1 <br />
   }, <br />
   { <br />
      "file":"/etc/remove 1", <br />
      "uid":"0", <br />
      "gid":"0", <br />
      "modtime":"1592179200", <br />
      "state":0 <br />
   }, <br />
   { <br />
      "file":"/etc/change 1", <br />
      "uid":"0", <br />
      "gid":"0", <br />
      "modtime":"1592185671", <br />
      "state":2 <br />
   } <br />
] <br />