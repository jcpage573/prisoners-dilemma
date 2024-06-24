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
	user, err := stripUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("POSTPRISREQ", user, "!!")
}

func (ward *Warden) GetPrisoner(w http.ResponseWriter, r *http.Request) {
	user, err := stripUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("GETPRISREQ", user, "!!")
}

func stripUser(r *http.Request) (string, error) {
	user := strings.TrimPrefix(r.URL.Path, "/user/")
	if user == "" || user == "/" {
		return "", fmt.Errorf("invalid or unspecified user '%s'", user)
	}
	return user, nil
}
