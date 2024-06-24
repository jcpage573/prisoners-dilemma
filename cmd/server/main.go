package main

import (
	"log"
	"net/http"

	"github.com/jcpage573/prisoners-dilemma/internal/server"
)

func main() {
	// New mux
	mux := http.NewServeMux()

	// Handler
	warden := server.NewWarden()

	// Register public routes
	mux.HandleFunc("/test", server.TestHandler)

	// Create a new mux for authenticated routes
	authMux := http.NewServeMux()

	// Register authenticated routes
	authMux.HandleFunc("POST /user/", warden.NewPrisoner)

	// Wrap the authMux in auth middleware
	authHandler := server.NewAuth(authMux)

	// Add the authenticated routes to the main mux
	mux.Handle("/", authHandler)

	// Wrap the entire mux in logger middleware
	server := server.Logger(mux)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", server))
}
