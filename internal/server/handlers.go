package server

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func TestHandler(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Erd Tree!")) }

type Warden struct {
	store net.Conn
}

func NewWarden() Warden {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	return Warden{store: conn}
}

func (ward *Warden) NewPrisoner(w http.ResponseWriter, r *http.Request) {
	// Extract the user part from the URL
	user := strings.TrimPrefix(r.URL.Path, "/user/")
	if user == "" || user == "/" {
		http.Error(w, "User not specified", http.StatusBadRequest)
		return
	}

	fmt.Println("PRISREQ", user, "!!")
}
