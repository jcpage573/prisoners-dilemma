package server

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

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
