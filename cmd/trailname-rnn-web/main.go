package main

import (
	"log"
	"net/http"
	"github.com/mburtless/trailname-rnn-web/pkg/routes"
	"github.com/mburtless/trailname-rnn-web/pkg/configs"
	"flag"
)

//var apiHostFlag string
//var portFlag string
//var envFlag string

/*var configVars = map[string]string {
	"APIHOST":		"localhost",
	"PORT":			"8000",
	"ENVIRONMENT":	"development",
}*/

type configSet struct {
	defVal	string
	envVar	string
	flagVar string
	help	string
	parsedVal	string
}

var configVars = map[string]*configSet{
	"apiHost":		{"localhost", "APIHOST", "apihost", "IP or hostname of the trailname-rnn API", ""},
	"port":			{"8000", "PORT", "port", "Port to listen on", ""},
	"environment":	{"development", "ENVIRONMENT", "environemnt", "Deployment environment", ""},
}

func init() {
	// Define our flags
	//flag.StringVar(configVars["apihost"]["parsedVal"], configVars["apihost"]["flagVar"], configVars["apihost"]["default"], configVars["apihost"]["help"])
	flag.StringVar(&configVars["apiHost"].parsedVal, configVars["apiHost"].flagVar, configVars["apiHost"].defVal, configVars["apiHost"].help)
	flag.StringVar(&configVars["port"].parsedVal, configVars["port"].flagVar, configVars["port"].defVal, configVars["port"].help)
	flag.StringVar(&configVars["environment"].parsedVal, configVars["environment"].flagVar, configVars["environment"].defVal, configVars["environment"].help)
	//flag.StringVar(&portFlag, "port", "8080", "Port to listen on")
	//flag.StringVar(&envFlag, "environment", "development", "Deployment environment")
}

func main() {
	// Parse flags
	flag.Parse()
	// If ParseEnv indicates APIHOST env doesn't exist, store user value or default passed via flag
	//if configs.ParseEnv("APIHOST") {
	//	log.Printf("APIHOST env found, ignoring flag value")
	//} else {
	//	// Send to configs for parsing into InstanceArgs
	//	log.Printf("APIHOST env not found, using flag value instead")
	//	configs.ParseFlags(&configVars["apiHost"].parsedVal)
	//}
	for _, val := range configVars {
		if configs.ParseEnv(val.envVar) {
			log.Printf("%s env found, ignoring flag value", val.envVar)
		} else {
			// Send to configs for parsing into InstanceArgs
			log.Printf("%s env not found, using flag value instead", val.envVar)
			configs.ParseFlags(val.envVar, &val.parsedVal)
		}
	}
	log.Printf("Starting %s environment on port %s; Targeting API host at %s", *configs.InstanceArgs["ENVIRONMENT"], *configs.InstanceArgs["PORT"], *configs.InstanceArgs["APIHOST"])
	// Init dynamic routes
	router := routes.NewRouter()
	// Set parent dir for static content
	staticDirectory := "./web"
	// Create routes for static content
	routes.StaticRouter(router, staticDirectory)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))

}
