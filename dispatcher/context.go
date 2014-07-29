package dispatcher

import (
	"fmt"
	"go.sancus.io/web/context"
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
