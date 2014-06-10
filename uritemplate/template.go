package uritemplate

import (
	"go.sancus.io/core/log"
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

// Literal
func (t *Template) appendLiteral(str string) {
	e := exprLiteral{literal: str}
	t.push(&e)
}

// EOL
func (t *Template) appendEOL() {
	e := exprSpecial{typ: exprEOL}
	t.push(&e)
}
