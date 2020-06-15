Overview
[Alt text](img/overview.JPG?raw=true "Overview")

使用說明

1. 使用superuser的權限執行install
$ sudo /bin/bash install

2. 修改監控目錄清單 monitor.list
$ vi monitor.list

範例
#------------------------
#ADD MONITORING DIRECTORY
#------------------------
/etc/
#/home/

3. 執行baseline,建立基準線
$ baseline

4. 執行monitor,監控程式會產生checkout.json
$ monitor


checkout.json 格式:

[
{
   "file": "<file path>" (str)
   "uid": "<uid>" (str)
   "gid": "<gid>" (str)
   "modtime": "<unix time(s)>" (str)
   "state": "< 0:delete , 1:add , 2:change >" (int)
},
…
]


checkout.json 範例:

[
   {
      "file":"/etc/add 1",
      "uid":"0",
      "gid":"0",
      "modtime":"1592185615",
      "state":1
   },
   {
      "file":"/etc/remove 1",
      "uid":"0",
      "gid":"0",
      "modtime":"1592179200",
      "state":0
   },
   {
      "file":"/etc/change 1",
      "uid":"0",
      "gid":"0",
      "modtime":"1592185671",
      "state":2
   }
]