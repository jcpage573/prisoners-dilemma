package server

import (
	"log"
	"net/http"
	"time"
)

// Logger is a middleware handler that does request logging
type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
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
