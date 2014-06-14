package uritemplate

import (
	"bytes"
	"go.sancus.io/core/log"
)

type exprSequence struct {
	expr []expression
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

// parser
func (e *exprSequence) addToken(t *token, p *parser) (bool, error) {
	switch t.typ {
	case tokenText:
		v := exprLiteral{literal: t.val}
		e.push(&v)
	case tokenEOL:
		v := exprSpecial{typ: exprEOL}
		e.push(&v)
	case tokenEOF:
		// we are done
		return false, nil
	case tokenLeftBrace:
		v := exprCapture{}
		p.stack.push(&v)
	case tokenLeftBracket:
		v := exprOptional{}
		p.stack.push(&v)
	case tokenError:
		return false, log.NewError("%s", t)
	default:
		return false, log.NewError("Sequence doesn't accept %s tokens", t)
	}

	return true, nil
}
