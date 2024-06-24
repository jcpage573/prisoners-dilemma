package main

import (
	"log"
	"net/http"

	"github.com/jcpage573/prisoners-dilemma/internal/server"
)

func main() {
	// New mux
	mux := http.NewServeMux()

	// Register routes here
	mux.HandleFunc("/test", server.TestHandler)
	mux.HandleFunc("/user/", server.NewPrisoner)

	// Wrap the mux in middleware
	server := server.Logger(mux)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", server))
}
