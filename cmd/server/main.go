package main

import (
	"log"
	"net/http"

	"github.com/jcpage573/prisoners-dilemma/internal/server"
)

func main() {
	http.HandleFunc("/test", server.TestHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
