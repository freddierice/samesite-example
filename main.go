package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// First host: first.trends.stream
	firstHost := r.Host("first.trends.stream").Subrouter()
	firstHost.HandleFunc("/", serveFirstPage).Methods("GET")
	firstHost.HandleFunc("/set", setCookie).Methods("GET")
	firstHost.HandleFunc("/number", getNumber).Methods("POST")

	// Second host: second.trends.stream
	secondHost := r.Host("second.trends.stream").Subrouter()
	secondHost.HandleFunc("/", serveSecondPage)

	// Default host: trends.stream
	defaultHost := r.Host("trends.stream").Subrouter()
	defaultHost.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to trends.stream")
	})

	// Set up CORS
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "https://second.trends.stream")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	// r.Use(corsMiddleware)
	firstHost.Use(corsMiddleware)

	// Start the server
	srv := &http.Server{
		Handler:      r,
		Addr:         ":443",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServeTLS("origin.pem", "origin.key"))
}
