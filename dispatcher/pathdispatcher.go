package dispatcher

import (
	"net/http"
)

type PathDispatcher struct{}

func (d *PathDispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
