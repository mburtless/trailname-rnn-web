package main

import (
	"log"
	"net/http"
	"github.com/mburtless/trailname-rnn-web/pkg/routes"
)


func main() {
	router := routes.NewRouter()
	//router.HandleFunc("/trailname/{startText}", handlers.GetTrailName).Methods("GET")
	staticDirectory := "./web/"
	routes.StaticRouter(router, staticDirectory)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", router))

}
