package dispatcher

import (
	"fmt"
	"go.sancus.io/web/context"
	"net/http"
)

// sets sancus.host and sancus.script_name headers
// required by the other dispatchers
func Prepare(m context.RequestContextMapper, r *http.Request) {
	c := GetArguments(m, r)

	// Host: -> Header["sancus.host"]
	if _, ok := c["HOST"].(string); !ok {
		c["HOST"] = r.Host
		r.URL.Host = ""
		r.Host = ""
	}

	// Header["sancus.script_name"]
	if _, ok := c["SCRIPT_NAME"].(string); !ok {
		c["SCRIPT_NAME"] = ""
	}
}

func Path(m context.RequestContextMapper, r *http.Request, path string, a ...interface{}) string {
	if len(a) > 0 {
		path = fmt.Sprintf(path, a)
	}

	if v, ok := GetArguments(m, r)["SCRIPT_NAME"]; ok {
		if str, ok := v.(string); ok && len(str) > 0 {
			path = str + path
		}
	}

	return path
}
