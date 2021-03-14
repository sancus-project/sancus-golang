package uritemplate

import (
	"bytes"
)

type exprOptional struct {
	exprSequence
}

func (e *exprOptional) Type() exprType {
	return exprOPTIONAL
}

func (e *exprOptional) String() string {
	var b bytes.Buffer
	b.WriteRune('[')
	for i, v := range e.expr {
		if i > 0 {
			b.WriteRune(' ')
		}
		b.WriteString(v.String())
	}
	b.WriteRune(']')
	return b.String()
}

// parser
//
func (e *exprOptional) addToken(t *token, p *parser) (bool, error) {
	if t.typ == tokenRightBracket {
		p.pop()
		return true, nil
	}

	return e.exprSequence.addToken(t, p)
}

func (e *exprOptional) addExpression(v expression, p *parser) (bool, error) {
	e.push(v)
	return true, nil
}
