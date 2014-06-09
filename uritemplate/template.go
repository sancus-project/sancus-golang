package uritemplate

import (
	"fmt"
	"go.sancus.io/core/log"
)

type exprType int

type expression interface {
	Type() exprType
	String() string

	addToken(t *token, s *parser) bool
}

const (
	exprLITERAL exprType = iota + 1
	exprEOL
	exprCAPTURE
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

// Capture
//
type exprCapture struct {
	key string
	options []expression
}

func (e *exprCapture) Type() exprType {
	return exprCAPTURE
}

func (e *exprCapture) String() string {
	return fmt.Sprintf("{%s}", log.NonEmptyString(e.key, "undefined"))
}

func (e *exprCapture) addToken(t *token, p *parser) bool {
	switch t.typ {
	case tokenIdentifier:
		if len(e.key) == 0 && len(t.val) > 0 {
			e.key = t.val
			return true
		}
	case tokenOption:
		s := exprLiteral{literal: t.val}
		e.options = append(e.options, &s)
		return true
	}
	return false
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

func (e *exprLiteral) addToken(t *token, p *parser) bool {
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

func (e *exprSpecial) addToken(t *token, p *parser) bool {
	return false
}

// EOL
func (t *Template) appendEOL() {
	e := exprSpecial{typ: exprEOL}
	t.append(&e)
}
