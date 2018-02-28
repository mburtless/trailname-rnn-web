package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/mburtless/trailname-rnn-web/pkg/handlers"
	"github.com/mburtless/trailname-rnn-web/pkg/logger"
	"log"
)

type Route struct {
	Name		string
	Method		string
	Pattern		string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
	/*Route {
		"GetTrailName",
		"POST",
		"/trailname/{starttext}",
		handlers.GetTestTrailName,
	},*/
	Route {
		"APIReq",
		"POST",
		"/api",
		handlers.ApiHandler,
	},
	Route {
		"Index",
		"GET",
		"/",
		handlers.IndexHandler,
	},
	Route {
		"Index",
		"GET",
		"/index.html",
		handlers.IndexHandler,
	},
}

var staticDirs = []string {"/css/", "/fonts/", "/js/", "/pages/"}

// NewRouter creates a mux router and adds routes for 
// all routes defined in the routes var
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		log.Printf("Creating %s route for %s", route.Name, route.Pattern)
		var handler http.Handler
		handler = route.HandlerFunc

		handler = logger.HandlerLog(handler, route.Name)
		router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)

	}
	return router
}

// StaticRouter creates a handler and adds routes for
// all static asset directories
func StaticRouter(router *mux.Router, staticDirectory string) {
	for _, dir := range staticDirs {
		pathValue := staticDirectory + dir
		log.Printf("Adding route to StaticHandler for %s", pathValue)
		handler := http.StripPrefix(dir, http.FileServer(http.Dir(pathValue)))
		handler = logger.HandlerLog(handler, "StaticHandler")
		router.PathPrefix(dir).Handler(handler)
	}
}
