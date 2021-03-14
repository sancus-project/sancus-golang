package uritemplate

import (
	"go.sancus.dev/sancus/log"
)

// parsing stack
type exprStack struct {
	exprSequence
}

func (s *exprStack) addToken(t *token, p *parser) (bool, error) {
	if e, ok := s.last().(container); ok {
		return e.addToken(t, p)
	}

	return false, log.NewError("Last expression in the stack (%s) doesn't accept subexpressions",
		s.last(), t)
}

// Parser
type parser struct {
	logger *log.Logger

	stack exprStack
	tmpl  *Template
}

func (p *parser) addToken(t token) (bool, error) {
	if p.stack.Len() == 0 {
		return p.tmpl.addToken(&t, p)
	} else if last, ok := p.stack.last().(container); ok {
		return last.addToken(&t, p)
	} else {
		return false, log.NewError("Unhandled token (%s) [last: %s]", t, p.stack.last())
	}
}

func (p *parser) pop() {
	e := p.stack.pop()
	if e == nil {
		p.logger.Panic("pop over empty stack")
	} else if p.stack.Len() == 0 {
		p.tmpl.push(e)
	} else if last, ok := p.stack.last().(container); ok {
		last.addExpression(e, p)
	} else {
		p.logger.Panic("parent expression doesn't support subexpressions")
	}
}

// Turn string into Template
func string2Template(str string, tmpl *Template) error {
	l := tmpl.logger
	// new parser
	p := &parser{logger: l, tmpl: tmpl}

	lexLogger := l.SubLogger(".lexer")
	lexLogger.Level = log.INFO

	lex := newLexer(str, lexLogger)

	for {
		cont, err := p.addToken(lex.nextToken())
		if cont {
			// eat all token
			continue
		}

		if err == nil {
			l.Trace("tmpl: %v", tmpl)
		}
		return err
	}

	return nil
}
