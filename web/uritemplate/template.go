package uritemplate

import (
	"go.sancus.dev/sancus/attic/log"
)

// Template
type Template struct {
	exprSequence

	logger *log.Logger
}

func NewTemplate(tmpl string, logger *log.Logger) (*Template, error) {
	t := &Template{logger: logger}
	err := string2Template(tmpl, t)

	if err != nil {
		return nil, err
	}
	return t, nil
}
