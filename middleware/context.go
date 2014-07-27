package middleware

import (
	"go.sancus.io/web/context"
	"net/http"
)

// ContextMiddleware removes the request from the map on exit
func RemoveContextMiddleware(ctx *context.ContextMap, h http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		defer ctx.RemoveAll(r)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
}
