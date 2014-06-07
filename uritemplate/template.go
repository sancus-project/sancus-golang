package uritemplate

import (
	"fmt"
	"go.sancus.io/core/log"
)

type exprType int

type expression interface {
	Type() exprType
	String() string
}

const (
	exprLiteral exprType = iota + 1
)

// Template
type Template struct {
	logger *log.Logger

	expr []expression
}

func NewTemplate(tmpl string, logger *log.Logger) (*Template, error) {
	t := &Template{logger: logger}
	err := string2Template(tmpl, t)

	if err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Template) append(e expression) {
	t.expr = append(t.expr, e)
}

// Literal string
//
type tmplLiteral struct {
	literal string
}

func (t *tmplLiteral) Type() exprType {
	return exprLiteral
}

func (t *tmplLiteral) String() string {
	return fmt.Sprintf("%q", t.literal)
}

func (t *Template) appendLiteral(str string) {
	e := tmplLiteral{literal: str}
	t.append(&e)
}
