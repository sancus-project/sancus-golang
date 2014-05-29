package uritemplate

import (
	"go.sancus.io/core/log"
)

type Template struct{}

func ParseTemplate(template string, l *log.Logger) *Template {
	lex := newLexer(template, l)
	for {
		t := lex.nextToken()
		l.Trace("ParseTemplate: t=%s", t)
		if t.typ == tokenEOF {
			break
		}
	}
	return nil
}
