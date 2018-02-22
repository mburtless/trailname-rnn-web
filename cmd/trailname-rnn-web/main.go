package main

import (
	"log"
	"net/http"
	"github.com/mburtless/trailname-rnn-web/pkg/routes"
	"github.com/mburtless/trailname-rnn-web/pkg/configs"
	"flag"
)

var apiHostFlag string

func init() {
	// Define our flags
	flag.StringVar(&apiHostFlag, "apihost", "localhost", "IP or hostname of the trailname-rnn API")
}

func main() {
	// Parse flags
	flag.Parse()
	// Send to configs for parsing into InstanceArgs
	configs.ParseFlags(&apiHostFlag)
	// Init dynamic routes
	router := routes.NewRouter()
	// Set parent dir for static content
	staticDirectory := "./web"
	// Create routes for static content
	routes.StaticRouter(router, staticDirectory)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))

}
