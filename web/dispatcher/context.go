package dispatcher

import (
	"fmt"
	"go.sancus.dev/sancus/web/context"
	"net/http"
)

var (
	DispatcherArgsLabel = "dispatcher-arguments"
)

type Arguments map[string]interface{}

//
func GetArguments(m context.RequestContextMapper, r *http.Request) Arguments {
	var args Arguments
	label := DispatcherArgsLabel

	v, ok := m.Get(r, label)
	if v == nil {
		args = make(Arguments)
		m.Set(r, label, args)
	} else if args, ok = v.(Arguments); !ok {
		panic(fmt.Sprintf("%s: already in use by incorrect data", label))
	}

	return args
}

func GetArgument(m context.RequestContextMapper, r *http.Request, key string) (interface{}, bool) {
	v, ok := GetArguments(m, r)[key]
	return v, ok
}

func HasArgument(m context.RequestContextMapper, r *http.Request, key string) bool {
	_, ok := GetArgument(m, r, key)
	return ok
}

func GetStringArgument(m context.RequestContextMapper, r *http.Request, key string, fallback string) string {
	v := GetArguments(m, r)[key]

	if str, ok := v.(string); ok {
		return str
	} else if str, ok := v.(fmt.Stringer); ok {
		return str.String()
	} else {
		return fallback
	}
}

func SetArgument(m context.RequestContextMapper, r *http.Request, key string, v interface{}) interface{} {
	args := GetArguments(m, r)
	v0 := args[key]
	args[key] = v
	return v0
}
