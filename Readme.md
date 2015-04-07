# zabbix-rest-bridge

Creates a REST bridge for a [Zabbix server](https://www.zabix.com).
The [Zabbix protocol](https://www.zabbix.com/documentation/2.2/manual/appendix/items/activepassive)
is _almost_ REST-like.  This application starts up a small REST server
handling POST and PUT methods and forwards to the given Zabbix server,
returning the response.  Now CURL can be used to populate Zabbix.

## Install

    go get github.com/dblezek/zabbix-rest-bridge

## Running

    zabbix-rest-bridge --server zabbix.example.com

By default, `zabbix-rest-bridge` listens on port `8987` on all interfaces, forwarding requests to `zabbix.example.com` on port `10051`.

## Getting Help

    zabbix-rest-bridge -h

## Examples

Values can be `POST`ed by specifying `POST` parameters.  `key` is the Zabbix key, `value` is the value to association, and `host` is the hostname where this value came from.  For instance to post 'foo=1234' from 'hostname'.

    curl -v -X POST -d host="hostname" -d key=foo -d value=1234 localhost:8987

Well formed Zabbix JSON can be posted.  Post the contents of TestData.txt

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

## License

The MIT License (MIT)

Copyright (c) 2015 Daniel Blezek

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
