package uritemplate

import (
	"go.sancus.io/core/log"
)

// parsing stack
type exprStack struct {
	stack []expression
}

func (s *exprStack) Len() int {
	return len(s.stack)
}

func (s *exprStack) append(e expression) {
	s.stack = append(s.stack, e)
}

func (s *exprStack) addToken(t token) bool {
	if l := len(s.stack) - 1; l >= 0 {
		e := s.stack[l]
		return e.addToken(t)
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
	stackLen := p.stack.Len()

	l.Trace("addToken: t=%s len=%v", t, stackLen)

	if stackLen == 0 {
		// nothing incomplete waiting

		switch t.typ {
		case tokenText:
			p.tmpl.appendLiteral(t.val)
		case tokenEOL:
			p.tmpl.appendEOL()
		case tokenEOF:
			// we are done
			return false
		default:
			p.logger.Panic("addToken: Unhandled token (%s)", t)
			return false

		}
	} else if !p.stack.addToken(t) {
		p.logger.Panic("addToken: Unhandled token (%s) [stackLen=%v]", t, stackLen)
		return false
	}

	// continue
	return true
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

	l.Trace("tmpl: %v", t.expr)
	return nil
}
