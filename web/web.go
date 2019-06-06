package web

import (
	"net/http"
)

// H converts a HandlerFunc into a Handler
func H(h http.HandlerFunc) http.Handler {
	return h
}
