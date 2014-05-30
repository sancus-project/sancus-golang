package dispatcher

import (
	"fmt"
	"net/http"
)

// sets sancus.host and sancus.script_name headers
// required by the other dispatchers
func Prepare(r *http.Request) {
	// Host: -> Header["sancus.host"]
	if v := r.Header["sancus.host"]; v == nil {
		r.Header["sancus.host"] = []string{r.Host}
		r.URL.Host = ""
		r.Host = ""
	}

	// Header["sancus.script_name"]
	if v := r.Header["sancus.script_name"]; v == nil {
		r.Header["sancus.script_name"] = []string{""}
	}
}

func Path(r *http.Request, path string, a ...interface{}) string {
	if len(a) > 0 {
		path = fmt.Sprintf(path, a)
	}

	if v := r.Header["sancus.script_name"]; v != nil && len(v[0]) > 0 {
		path = v[0] + path
	}

	return path
}
