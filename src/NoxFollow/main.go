package main

import (
    "github.com/NyroCodes/NoxFollow/api"
    "github.com/NyroCodes/NoxFollow/db"
    "log"
    "net/http"
)

func logRequestHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		// call the original http.Handler we're wrapping
		h.ServeHTTP(w, r)

		// gather information about request and log it
		uri := r.URL.String()

		log.Println(uri)
	}

	// http.HandlerFunc wraps a function so that it
	// implements http.Handler interface
	return http.HandlerFunc(fn)
}

func main() {
	
	// Init Redis Connection
	db.Init("redis:6379")
	
	mux := http.NewServeMux()
	
	redirect := &api.RedirectHandler{}
	redirect.Initialize(mux)
	
	follow := &api.FollowHandler{}
	follow.Initialize(mux)
	
	logMux := logRequestHandler(mux)

	log.Println("Listen at :8080")
	log.Fatal(http.ListenAndServe(":8080", logMux))
}
