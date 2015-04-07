### zabbix-rest-bridge

Creates a REST bridge for a Zabbix server.  The Zabbix protocol (https://www.zabbix.com/documentation/2.2/manual/appendix/items/activepassive) is _almost_ REST-like.  This application starts up a small REST server handling POST and GET methods and forwards to the given Zabbix server, returning the response.  Now CURL can be used to populate Zabbix.

Examples:

Post 'key=1234' from 'hostname':

  curl -v -X POST -d host="hostname" -d key=foo -d value=1234 localhost:8987

Post the contents of TestData.txt

  curl -v -X PUT -d @TestData.txt localhost:8987

TestData.txt is:
{
   "request":"agent data",
   "data":[
       {
           "host":"hostname",
           "key":"foo",
           "value":"1234"
       }
   ]
}
