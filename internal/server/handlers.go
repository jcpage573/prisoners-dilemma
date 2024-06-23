package server

import (
	"fmt"
	"net/http"
	"strings"
)

func TestHandler(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Erd Tree!")) }

// NewPrisoner handles requests to /user/someuser
func NewPrisoner(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the user part from the URL
	user := strings.TrimPrefix(r.URL.Path, "/user/")
	if user == "" || user == "/" {
		http.Error(w, "User not specified", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "New prisoner created for user: %s", user)
}
