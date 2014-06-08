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
	exprLITERAL exprType = iota + 1
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
type exprLiteral struct {
	literal string
}

func (e *exprLiteral) Type() exprType {
	return exprLITERAL
}

func (e *exprLiteral) String() string {
	return fmt.Sprintf("%q", e.literal)
}

func (e *exprLiteral) addToken(t token) bool {
	return false
}

func (t *Template) appendLiteral(str string) {
	e := exprLiteral{literal: str}
	t.append(&e)
}

// Special logic elements
//
type exprSpecial struct {
	typ exprType
}

func (e *exprSpecial) Type() exprType {
	return e.typ
}

func (e *exprSpecial) String() string {
	switch e.typ {
	case exprEOL:
		return "EOL"
	default:
		return "unknown"
	}
}

func (e *exprSpecial) addToken(t token) bool {
	return false
}

// EOL
func (t *Template) appendEOL() {
	e := exprSpecial{typ: exprEOL}
	t.append(&e)
}
