package middleware

import (
	"go.sancus.io/web/context"
	"go.sancus.io/web/dispatcher"
	"net/http"
)

// RemoveTrailingSlash redirects to the same URL but without the trailing /
func RemoveTrailingSlash(m context.RequestContextMapper, h http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if last := len(path) - 1; last > 0 && path[last] == '/' {
			http.Redirect(w, r, dispatcher.Path(m, r, path[:last]), 301)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
}
