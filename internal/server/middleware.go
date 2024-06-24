package server

import (
	"log"
	"net/http"
	"time"
)

// Logger is a middleware handler that does request logging
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}

// Auth is a middleware handler that verifies user requestor identities
const keyHeaderName string = "X-PRISONER-KEY"

type Auth struct {
	handler http.Handler
}

func (a *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get(keyHeaderName)
	if apiKey == "" {
		http.Error(w, "API Key required", http.StatusUnauthorized)
	}

	a.handler.ServeHTTP(w, r)
}

func NewAuth(handlerToWrap http.Handler) *Auth {
	return &Auth{handlerToWrap}
}
