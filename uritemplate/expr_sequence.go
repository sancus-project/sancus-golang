package uritemplate

import (
	"bytes"
)

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
	if l := len(e.expr) - 1; l >= 0 {
		v := e.expr[l]
		e.expr = e.expr[:l]
		return v
	}
	return nil
}

func (e *exprSequence) push(v expression) {
	e.expr = append(e.expr, v)
}
