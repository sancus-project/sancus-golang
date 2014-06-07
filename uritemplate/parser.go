package uritemplate

import (
	"go.sancus.io/core/log"
)

// Turn string into Template
func string2Template(str string, tmpl *Template) error {
	l := tmpl.logger

	lexLogger := l.SubLogger(".lexer")
	lexLogger.Level = log.INFO

	lex := newLexer(str, lexLogger)

	for {
		t := lex.nextToken()
		l.Trace("ParseTemplate: %s", t)
		if t.typ == tokenEOF {
			break
		}
	}
	return nil
}
