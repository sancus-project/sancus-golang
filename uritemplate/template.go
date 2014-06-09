package uritemplate

import (
	"bytes"
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
	exprSEQUENCE
	exprCAPTURE
)

// Sequence
type exprSequence struct {
	expr []expression
}

func (e *exprSequence) Type() exprType {
	return exprSEQUENCE
}

func (e *exprSequence) String() string {
	var b bytes.Buffer
	for i, v := range e.expr {
		if i > 0 {
			b.WriteRune(' ')
		}
		b.WriteString(v.String())
	}
	return b.String()
}

func (e *exprSequence) Len() int {
	return len(e.expr)
}

func (e *exprSequence) last() expression {
	if l := len(e.expr) - 1; l >= 0 {
		return e.expr[l]
	}
	return nil
}

func (e *exprSequence) pop() expression {
	if l := len(e.expr) -1; l >= 0 {
		v := e.expr[l]
		e.expr = e.expr[:l]
		return v
	}
	return nil
}

func (e *exprSequence) push(v expression) {
	e.expr = append(e.expr, v)
}

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

// Capture
//
type exprCapture struct {
	key     string
	options []expression
}

func (e *exprCapture) Type() exprType {
	return exprCAPTURE
}

func (e *exprCapture) String() string {
	return fmt.Sprintf("{%s}", log.NonEmptyString(e.key, "undefined"))
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

func (t *Template) appendLiteral(str string) {
	e := exprLiteral{literal: str}
	t.push(&e)
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

// EOL
func (t *Template) appendEOL() {
	e := exprSpecial{typ: exprEOL}
	t.push(&e)
}
