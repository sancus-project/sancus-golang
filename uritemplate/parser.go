package uritemplate

import (
	"go.sancus.io/core/log"
)

type Template struct{}

// Turn string into Template
func ParseTemplate(template string, l *log.Logger) *Template {
	lexLogger := l.SubLogger(".lexer")
	lexLogger.Level = log.INFO
	lex := newLexer(template, lexLogger)

	for {
		t := lex.nextToken()
		l.Trace("ParseTemplate: %s", t)
		if t.typ == tokenEOF {
			break
		}
	}
	return nil
}
