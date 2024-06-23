package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Erd Tree!")) })
	log.Fatal(http.ListenAndServe(":8080", nil))
}
