package dispatcher

import (
	"go.sancus.io/core/log"
	"go.sancus.io/web/context"
	"go.sancus.io/web/uritemplate"
	"net/http"
)

// PathDispatcher
type PathDispatcher struct {
	ContextMap *context.ContextMap
	Logger     *log.Logger
}

func NewPathDispatcher(ctx *context.ContextMap, loggerName string) *PathDispatcher {
	return &PathDispatcher{
		ContextMap: ctx,
		Logger:     log.GetLogger(loggerName),
	}
}

func (d *PathDispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.Logger.Debug("%s %s", r.Method, r.URL.Path)

	http.NotFound(w, r)
}

func (d *PathDispatcher) AddHandler(name string, pattern string, handler http.Handler) {
	l := d.Logger.SubLogger(".%s", log.NonEmptyString(name, "unnamed"))
	l.Level = log.DEBUG

	uritemplate.NewTemplate(pattern, l)
}
