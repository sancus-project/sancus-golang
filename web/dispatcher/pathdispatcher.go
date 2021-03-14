package dispatcher

import (
	"go.sancus.dev/sancus/attic/log"
	"go.sancus.dev/sancus/web/context"
	"go.sancus.dev/sancus/web/uritemplate"
	"net/http"
)

// PathDispatcher
type PathDispatcher struct {
	ContextMap context.RequestContextMapper
	Logger     *log.Logger
}

func NewPathDispatcher(m context.RequestContextMapper, loggerName string) *PathDispatcher {
	return &PathDispatcher{
		ContextMap: m,
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
