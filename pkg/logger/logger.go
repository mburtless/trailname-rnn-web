package logger

import (
	"log"
	"net/http"
	"time"
)

// Logger logs activity of a passed http.Handler to console.  It does so
// by acting as a wrapper around the passed handler and returning a new
// anonymous http.Handler function that performs the serving and logging
func HandlerLog(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Begin serving the passed handler
		inner.ServeHTTP(w, r)

		// Log the activity to console
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
