package uritemplate

import (
	"go.sancus.io/core/log"
)

// parsing stack
type exprStack struct {
	exprSequence
}

func (s *exprStack) addToken(t *token, p *parser) bool {
	if e, ok := s.last().(container); ok {
		return e.addToken(t, p)
	}
	p.logger.Error("addToken: Last expression in the stack (%s) doesn't accept subexpressions",
		s.last(), t)
	return false
}

// Parser
type parser struct {
	logger *log.Logger

	stack exprStack
	tmpl  *Template
}

func (p *parser) addToken(t token) bool {
	if p.stack.Len() == 0 {
		return p.tmpl.addToken(&t, p)
	} else if last, ok := p.stack.last().(container); ok {
		return last.addToken(&t, p)
	} else {
		p.logger.Error("addToken: Unhandled token (%s) [last: %s]", t, p.stack.last())
		return false
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
	t := &Template{}
	p := &parser{logger: l, tmpl: t}

	lexLogger := l.SubLogger(".lexer")
	lexLogger.Level = log.INFO

	lex := newLexer(str, lexLogger)

	for p.addToken(lex.nextToken()) {
		// eat all tokens
	}

	l.Trace("tmpl: %v", t)
	return nil
}
