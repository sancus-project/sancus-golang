package dispatcher

import (
	"go.sancus.io/core/log"
	"go.sancus.io/web/context"
	"net/http"
	"regexp"
)

type Bridge struct {
	Logger     *log.Logger
	ContextMap context.RequestContextMapper

	scriptName *regexp.Regexp
	handler    http.Handler
}

func NewBridge(m context.RequestContextMapper, scriptName string, logger string, handler http.Handler) *Bridge {
	var pattern string

	l := log.GetLogger(logger)
	d := &Bridge{
		Logger:     l,
		ContextMap: m,
	}

	if scriptName == "" || scriptName == "/" {
		pattern = "^()(/.*)$"
	} else {
		pattern = "^(" + scriptName + ")(/.*)$"
	}

	l.Debug("NewBridge(%q) -> %q", scriptName, pattern)

	d.scriptName = regexp.MustCompile(pattern)
	d.handler = handler
	return d
}

/*
 * Strips scriptName out of URL.Path
 */
func (d *Bridge) stripScriptName(r *http.Request) bool {
	l := d.Logger

	Prepare(d.ContextMap, r)

	if m := d.scriptName.FindStringSubmatch(r.URL.Path); m != nil {
		// script_name found
		if m[1] != "" {
			script_name := r.Header["sancus.script_name"][0] + m[1]
			r.Header["sancus.script_name"] = []string{script_name}
			r.URL.Path = m[2]

			if !l.Debug("%s%s -> %s (%q)", script_name, r.URL.Path, r.URL.Path, m) {
				l.Verbose("%s%s -> %s", script_name, r.URL.Path, r.URL.Path)
			}
		}
		return true
	}

	l.Warn("%s: No ScriptName match!", r.URL.Path)
	return false
}

// Regular net.http.Handler to strip scriptName off
func (d *Bridge) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if d.stripScriptName(r) {
		d.handler.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}
