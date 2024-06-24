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

	fmt.Println("PRISREQ", user, "!!")
}
