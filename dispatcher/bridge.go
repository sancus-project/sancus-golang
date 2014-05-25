package dispatcher

import (
	"go.sancus.io/core/log"
	"net/http"
	"regexp"
)

type Bridge struct {
	Logger *log.Logger

	scriptName *regexp.Regexp
	handler    http.Handler
}

func NewBridge(scriptName string, logger string, handler http.Handler) *Bridge {
	var pattern string

	l := log.GetLogger(logger)
	d := &Bridge{Logger: l}

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
	var script_name string
	l := d.Logger

	Prepare(r)
	script_name = r.Header["sancus.script_name"][0]

	m := d.scriptName.FindStringSubmatch(r.URL.Path)
	if m != nil {
		// script_name found
		script_name += m[1]

		r.URL.Path = m[2]
		r.Header["sancus.script_name"] = []string{script_name}

		if l.IsLoggable(log.DEBUG) {
			l.Debug("%s%s -> %s (%q)", script_name, r.URL.Path, r.URL.Path, m)
		} else {
			l.Verbose("%s%s -> %s", script_name, r.URL.Path, r.URL.Path)
		}

		return true
	} else if l.IsLoggable(log.DEBUG) {
		l.Warn("%s: No ScriptName match! (%q)", r.URL.Path, m)
	} else {
		l.Warn("%s: No ScriptName match!", r.URL.Path)
	}
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
