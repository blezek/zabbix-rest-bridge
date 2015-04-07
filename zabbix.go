package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/codegangsta/cli"
	"github.com/gorilla/mux"
	"github.com/sfreiberg/zbxutils"
)

type Body struct {
	Host  string
	Key   string
	Value string
}

func main() {

	app := cli.NewApp()
	app.Name = "zabbix-rest-bridge"
	app.Usage = usage
	app.Version = "1.0.0"
	app.Authors = []cli.Author{cli.Author{"Daniel Blezek", "daniel.blezek@gmail.com"}}
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "timeout,t",
			Value: 5,
			Usage: "Zabbix request timeout in seconds",
		},
		cli.IntFlag{
			Name:  "zabbix-port,xz",
			Value: 10051,
			Usage: "Remote Zabbix port",
		},
		cli.StringFlag{
			Name:  "server,s",
			Usage: "Remote Zabbix server name",
		},
		cli.IntFlag{
			Name:  "port,p",
			Value: 8987,
			Usage: "Server port",
		},
		cli.StringFlag{
			Name:  "interface,i",
			Value: "",
			Usage: "Interface to listen on, default is all, 127.0.0.1 would be localhost only",
		},
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Verbose logging",
		},
	}
	app.Action = func(context *cli.Context) {
		verbose := context.Bool("verbose")
		timeout := time.Duration(context.Int("timeout")) * time.Second
		server := zbxutils.NewAgentHostPort(context.String("server"), context.Int("zabbix-port"))
		r := mux.NewRouter()
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			v := Body{r.FormValue("host"), r.FormValue("key"), r.FormValue("value")}
			buffer := &bytes.Buffer{}
			bodyTemplate, err := template.New("template").Parse(templateString)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error parsing template: %v", err), 200)
				return
			}
			err = bodyTemplate.Execute(buffer, v)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error executing template: %v", err), 200)
				return
			}

			if verbose {
				log.Printf("Sending %v to Zabbix\n", buffer.String())
			}
			response, err := server.GetWithTimeout(buffer.String(), timeout)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error from Zabbix server: %v", err), 200)
				return
			}
			if verbose {
				log.Printf("Got %v from Zabbix\n", string(response.Data))
			}
			w.Write(response.Data)
			return
		}).Methods("POST")
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			buffer, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, fmt.Sprintf("%v", err), 200)
				return
			}
			if verbose {
				log.Printf("Sending %v to Zabbix\n", string(buffer))
			}
			response, err := server.GetWithTimeout(string(buffer), timeout)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error from Zabbix server: %v", err), 200)
				return
			}
			if verbose {
				log.Printf("Got %v from Zabbix\n", string(response.Data))
			}
			w.Write(response.Data)
			return
		}).Methods("PUT")
		log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", context.String("interface"), context.Int("port")), r))
	}

	app.Run(os.Args)
}

var usage = `Creates a REST bridge for a Zabbix server.  The Zabbix
protocol (*) is _almost_ REST-like.  This application starts up a
small REST server handling POST and GET methods and forwards to the
given Zabbix server, returning the response.  Now CURL can be used to
populate Zabbix.

(*)https://www.zabbix.com/documentation/2.2/manual/appendix/items/activepassive

Examples:

Post 'foo=1234' from 'hostname':

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
`

var templateString = `{
   "request":"agent data",
   "data":[
       {
           "host":"{{.Host}}",
           "key":"{{.Key}}",
           "value":"{{.Value}}"
       }
   ]
}`
