package server

import (
	"fmt"
	"net/http"
	"strings"
)

func TestHandler(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Erd Tree!")) }

func NewPrisoner(w http.ResponseWriter, r *http.Request) {
	// Extract the user part from the URL
	user := strings.TrimPrefix(r.URL.Path, "/user/")
	if user == "" || user == "/" {
		http.Error(w, "User not specified", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		handleGetPrisoner(w, r, user)
	case http.MethodPost:
		handlePostPrisoner(w, r, user)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetPrisoner(w http.ResponseWriter, r *http.Request, user string) {
	// Here you would typically fetch the prisoner's data
	// For now, we'll just return a placeholder message
	fmt.Fprintf(w, "Prisoner data for user: %s", user)
}

func handlePostPrisoner(w http.ResponseWriter, r *http.Request, user string) {
	// This is your existing logic for creating a new prisoner
	fmt.Fprintf(w, "New prisoner created for user: %s", user)
}
