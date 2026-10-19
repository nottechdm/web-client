package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	// endpoint for the root path
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello, World!"))
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	})
	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 2 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	// start the server

	log.Println("Server listening on :8080")
	log.Fatal(server.ListenAndServe())
}
