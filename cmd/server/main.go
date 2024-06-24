package main

import (
	"log"
	"net"
	"net/http"

	"github.com/jcpage573/prisoners-dilemma/internal/server"
)

func main() {
	// New mux
	mux := http.NewServeMux()

	// Connect to the datastore
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}

	// Handler
	warden := server.NewWarden(conn)

	// Register public routes
	mux.HandleFunc("/test", server.TestHandler)
	mux.HandleFunc("POST /user/", warden.NewPrisoner)

	// Create a new mux for authenticated routes
	authMux := http.NewServeMux()

	// Register authenticated routes
	authMux.HandleFunc("GET /user/", warden.GetPrisoner)

	// Wrap the authMux in auth middleware
	authHandler := server.NewAuth(authMux, conn)

	// Add the authenticated routes to the main mux
	mux.Handle("/", authHandler)

	// Wrap the entire mux in logger middleware
	server := server.Logger(mux)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", server))
}
