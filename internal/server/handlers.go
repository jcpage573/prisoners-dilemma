package server

import (
	"fmt"
	"net/http"
	"strings"
)

func TestHandler(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Erd Tree!")) }

// NewPrisoner handles requests to /user/someuser
func NewPrisoner(w http.ResponseWriter, r *http.Request) {
	// Extract the user part from the URL
	user := strings.TrimPrefix(r.URL.Path, "/user/")
	if user == "" || user == "/" {
		http.Error(w, "User not specified", http.StatusBadRequest)
		return
	}

	// Process the request with the extracted user
	fmt.Fprintf(w, "New prisoner created for user: %s", user)
}
