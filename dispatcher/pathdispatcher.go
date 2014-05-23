package dispatcher

import (
	"go.sancus.io/core/log"
	"net/http"
)

// PathDispatcher
type PathDispatcher struct{}

func (d *PathDispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Info("PathDispatcher", "%v", r)
	http.NotFound(w, r)
}

func (d *PathDispatcher) AddNamed(name string, pattern string, handler http.Handler) {
}

func (d *PathDispatcher) Add(pattern string, handler http.Handler) {
	d.AddNamed("", pattern, handler)
}

func (d *PathDispatcher) AddFuncNamed(name string, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	d.AddNamed(name, pattern, http.HandlerFunc(handler))
}

func (d *PathDispatcher) AddFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	d.AddNamed("", pattern, http.HandlerFunc(handler))
}
