package uritemplate

import (
	"go.sancus.io/core/log"
)

// Parser
type parser struct {
	logger *log.Logger

	stack []expression
	tmpl  *Template
}

func (p *parser) addToken(t token) bool {
	l := p.logger
	stackLen := len(p.stack)

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

		// continue
		return true
	}

	p.logger.Panic("addToken: Unhandled token (%s)", t)
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

	l.Trace("tmpl: %v", t.expr)
	return nil
}
