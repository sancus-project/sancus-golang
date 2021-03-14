package middleware

import (
	"go.sancus.dev/sancus/web/context"
	"go.sancus.dev/sancus/web/dispatcher"
	"net/http"
	"regexp"
)

func ScriptName(ctx context.RequestContextMapper, scriptName string, h http.Handler) http.Handler {
	if scriptName == "" || scriptName == "/" {
		scriptName = "^()(/.*)$"
	} else {
		scriptName = "^(" + scriptName + ")(/.*)$"
	}

	pattern := regexp.MustCompile(scriptName)
	f := func(w http.ResponseWriter, r *http.Request) {
		dispatcher.Prepare(ctx, r)
		if m := pattern.FindStringSubmatch(r.URL.Path); m != nil {
			if m[1] != "" {
				str := dispatcher.GetStringArgument(ctx, r, "SCRIPT_NAME", "") + m[1]
				dispatcher.SetArgument(ctx, r, "SCRIPT_NAME", str)
				r.URL.Path = m[2]
			}
			h.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	}
	return http.HandlerFunc(f)
}
