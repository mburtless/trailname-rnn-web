package main

import (
	"log"
	"net/http"
	"github.com/mburtless/trailname-rnn-web/pkg/routes"
	"github.com/mburtless/trailname-rnn-web/pkg/configs"
	"flag"
)

func init() {
	// Define our flags
	flag.StringVar(&configs.ConfigVars["apiHost"].ParsedVal, configs.ConfigVars["apiHost"].FlagVar, configs.ConfigVars["apiHost"].DefVal, configs.ConfigVars["apiHost"].Help)
	flag.StringVar(&configs.ConfigVars["port"].ParsedVal, configs.ConfigVars["port"].FlagVar, configs.ConfigVars["port"].DefVal, configs.ConfigVars["port"].Help)
	flag.StringVar(&configs.ConfigVars["environment"].ParsedVal, configs.ConfigVars["environment"].FlagVar, configs.ConfigVars["environment"].DefVal, configs.ConfigVars["environment"].Help)
}

func main() {
	// Parse flags
	flag.Parse()
	// Parse all env vars and flags as needed
	configs.ParseConfigVars()
	log.Printf("Starting %s environment on port %s; Targeting API host at %s", configs.ConfigVars["environment"].ParsedVal, configs.ConfigVars["port"].ParsedVal, configs.ConfigVars["apiHost"].ParsedVal)
	// Init dynamic routes
	router := routes.NewRouter()
	// Set parent dir for static content
	staticDirectory := "./web"
	// Create routes for static content
	routes.StaticRouter(router, staticDirectory)
	//routes.IndexRouter(router, staticDirectory + "/pages/index.html")
	log.Fatal(http.ListenAndServe("0.0.0.0:" + configs.ConfigVars["port"].ParsedVal, router))

}
