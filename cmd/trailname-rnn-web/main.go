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



func init() {
	// Define our flags
	//flag.StringVar(configVars["apihost"]["ParsedVal"], configVars["apihost"]["FlagVar"], configVars["apihost"]["default"], configVars["apihost"]["Help"])
	flag.StringVar(&configs.ConfigVars["apiHost"].ParsedVal, configs.ConfigVars["apiHost"].FlagVar, configs.ConfigVars["apiHost"].DefVal, configs.ConfigVars["apiHost"].Help)
	flag.StringVar(&configs.ConfigVars["port"].ParsedVal, configs.ConfigVars["port"].FlagVar, configs.ConfigVars["port"].DefVal, configs.ConfigVars["port"].Help)
	flag.StringVar(&configs.ConfigVars["environment"].ParsedVal, configs.ConfigVars["environment"].FlagVar, configs.ConfigVars["environment"].DefVal, configs.ConfigVars["environment"].Help)
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
	//	configs.ParseFlags(&ConfigVars["apiHost"].ParsedVal)
	//}
	for _, val := range configs.ConfigVars {
		if configs.ParseEnv(val.EnvVar) {
			log.Printf("%s env found, ignoring flag value", val.EnvVar)
		} else {
			// Send to configs for parsing into InstanceArgs
			log.Printf("%s env not found, using flag value instead", val.EnvVar)
			configs.ParseFlags(val.EnvVar, &val.ParsedVal)
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
