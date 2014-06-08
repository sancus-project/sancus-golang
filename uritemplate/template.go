package uritemplate

import (
	"fmt"
	"go.sancus.io/core/log"
)

type exprType int

type expression interface {
	Type() exprType
	String() string

	addToken(t token) bool
}

const (
	exprLiteral exprType = iota + 1
	exprEOL
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

func (t *tmplLiteral) addToken(t token) bool {
	return false
}

func (t *Template) appendLiteral(str string) {
	e := tmplLiteral{literal: str}
	t.append(&e)
}

// Special logic elements
//
type tmplSpecial struct {
	typ exprType
}

func (t *tmplSpecial) Type() exprType {
	return t.typ
}

func (t *tmplSpecial) String() string {
	switch t.typ {
	case exprEOL:
		return "EOL"
	default:
		return "unknown"
	}
}

func (t *tmplSpecial) addToken(t token) bool {
	return false
}

// EOL
func (t *Template) appendEOL() {
	e := tmplSpecial{typ: exprEOL}
	t.append(&e)
}
