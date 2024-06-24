package server

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/jcpage573/prisoners-dilemma/internal/storage"
)

func TestHandler(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Erd Tree!")) }

type Warden struct {
	store store
}

type store struct {
	conn   net.Conn
	reader *storage.Reader
}

func NewWarden() Warden {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	return Warden{store: store{conn, storage.NewReader(conn)}}
}

func (ward *Warden) NewPrisoner(w http.ResponseWriter, r *http.Request) {
	user, err := stripUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	code, err := storage.NewCommand("SET", user, "hash(radagon)").Execute(ward.store.conn)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create new user: %s (code %d)", err.Error(), code), http.StatusInternalServerError)
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
