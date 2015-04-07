# zabbix-rest-bridge

Creates a REST bridge for a [Zabbix server](https://www.zabix.com).  The [Zabbix protocol](https://www.zabbix.com/documentation/2.2/manual/appendix/items/activepassive) is _almost_ REST-like.  This application starts up a small REST server handling POST and GET methods and forwards to the given Zabbix server, returning the response.  Now CURL can be used to populate Zabbix.

## Install

    go get github.com/dblezek/zabbix-rest-bridge

## Running

    zabbix-rest-bridge --server zabbix.example.com


## Examples

Post 'key=1234' from 'hostname':

    curl -v -X POST -d host="hostname" -d key=foo -d value=1234 localhost:8987

Post the contents of TestData.txt

    curl -v -X PUT -d @TestData.txt localhost:8987

TestData.txt:
```
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
```

## Special thanks

- https://github.com/codegangsta/cli -- Command line processing extraordinaire
- https://github.com/sfreiberg/zbxutils -- For doing all the Zabbix heavy lifting
- https://github.com/gorilla/mux -- Best HTTP Router Ever
