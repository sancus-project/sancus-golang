package uritemplate

import (
	"go.sancus.io/core/log"
)

// parsing stack
type exprStack struct {
	exprSequence
}

func (s *exprStack) addToken(t *token, p *parser) bool {
	if e := s.last(); e != nil {
		return e.addToken(t, p)
	}
	return false
}

// Parser
type parser struct {
	logger *log.Logger

	stack exprStack
	tmpl  *Template
}

func (p *parser) addToken(t token) bool {
	l := p.logger
	last := p.stack.last()

	if last == nil {
		// nothing incomplete waiting
		l.Trace("addToken: t=%s", t)

		switch t.typ {
		case tokenText:
			p.tmpl.appendLiteral(t.val)
		case tokenEOL:
			p.tmpl.appendEOL()
		case tokenEOF:
			// we are done
			return false
		case tokenLeftBrace:
			p.startCapture()
		default:
			p.logger.Panic("addToken: Unhandled token (%s)", t)
			return false

		}
	} else if !last.addToken(&t, p) {
		p.logger.Panic("addToken: Unhandled token (%s) [last: %s]", t, last)
		return false
	}

	// continue
	return true
}

func (p *parser) startCapture() {
	e := exprCapture{}
	p.stack.push(&e)
}

func (p *parser) pop() {
	e := p.stack.pop()
	if e == nil {
		p.logger.Panic("pop over empty stack")
	} else if p.stack.Len() == 0 {
		p.tmpl.push(e)
	} else {
		p.logger.Panic("pop: multilevel not yet implemented")
	}
}

// expr.addToken()
func (e *exprLiteral) addToken(t *token, p *parser) bool {
	return false
}

func (e *exprSpecial) addToken(t *token, p *parser) bool {
	return false
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
	case tokenRightBrace:
		p.pop()
		return true
	}
	return false
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
