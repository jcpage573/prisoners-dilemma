package server

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
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

	// Generate a new API key
	apiKey, err := generateAPIKey()
	if err != nil {
		http.Error(w, "Failed to generate API key", http.StatusInternalServerError)
		return
	}

	// Store the hashed API key with the hashed user as the key
	code, err := storage.NewCommand("SET", hashString(apiKey), user).Execute(ward.store.conn)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create new user: %s (code %d)", err.Error(), code), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Created new prisoner '%s'\nYour API key is %s\nSTORE THIS SOMEWHERE SAFE", user, apiKey)
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

func generateAPIKey() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func hashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
