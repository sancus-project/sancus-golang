package uritemplate

import (
	"fmt"
	"go.sancus.io/core/log"
)

type exprCapture struct {
	key     string
	options []expression
}

func (e *exprCapture) Type() exprType {
	return exprCAPTURE
}

func (e *exprCapture) String() string {
	return fmt.Sprintf("{%s}", log.NonEmptyString(e.key, "undefined"))
}

// lexer
func lexCaptureID(l *lexer) stateFn {
	var r rune
	l.log.Trace("state:lexCaptureID, %v..%v/%v",
		l.start, l.pos, len(l.input))

	if r = l.next(); (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		l.log.Debug("state:lexCaptureID r:%q (%v) [%v..%v/%v] ",
			r, l.width,
			l.start, l.pos, len(l.input))

		for {
			var emit *token
			var next stateFn
			r = l.next()

			l.log.Debug("state:lexText r:%q (%v) [%v..%v/%v] ",
				r, l.width,
				l.start, l.pos, len(l.input))

			l.log.Debug("state:lexCaptureID r:%q (%v) [%v..%v/%v] ",
				r, l.width,
				l.start, l.pos, len(l.input))

			if (r >= 'a' && r <= 'z') ||
				(r >= 'A' && r <= 'Z') ||
				(r >= '0' && r <= '9') ||
				r == '_' {
				// collecting identifier
				continue
			} else if r == '}' {
				// standard capture
				emit = &token{typ: tokenRightBrace}
				next = lexText
			} else if r == ':' {
				// capture with defined options
				next = lexCaptureOption
			} else {
				break
			}

			l.backup()
			l.emit(tokenIdentifier)
			l.restore()

			if emit != nil {
				l.emit(emit.typ)
			} else {
				l.ignore()
			}
			return next
		}

		l.emit(tokenIdentifier)
		return l.fail("incomplete capture")
	}

	return l.fail("invalid or missing missing capture ID")
}

func lexCaptureOption(l *lexer) stateFn {
	l.log.Trace("state:lexCaptureOption, %v..%v/%v",
		l.start, l.pos, len(l.input))

	for {
		var emit *token
		var next stateFn
		r := l.next()

		l.log.Debug("state:lexCaptureOption r:%q (%v) [%v..%v/%v] ",
			r, l.width,
			l.start, l.pos, len(l.input))

		if r == eof {
			break
		} else if r == '}' {
			emit = &token{typ: tokenRightBrace}
			next = lexText
		} else if r == '|' {
			next = lexCaptureOption
		} else {
			continue
		}

		l.backup()
		l.emit(tokenOption)
		l.restore()

		if emit != nil {
			l.emit(emit.typ)
		} else {
			l.ignore()
		}
		return next
	}

	if l.pos > l.start {
		l.emit(tokenOption)
	}
	return l.fail("incomplete capture")
}

// parser
func (e *exprCapture) addToken(t *token, p *parser) (bool, error) {
	switch t.typ {
	case tokenIdentifier:
		if len(e.key) == 0 && len(t.val) > 0 {
			e.key = t.val
		} else {
			s := "Capture: Identifier not accepted"
			p.logger.Panic(s)

			// this should not be reached, but well... fail.
			return false, log.NewError(s)
		}
	case tokenOption:
		s := exprLiteral{literal: t.val}
		e.options = append(e.options, &s)
	case tokenRightBrace:
		p.pop()
	case tokenError:
		return false, log.NewError("%s", t)
	default:
		return false, log.NewError("Capture doesn't accept %s tokens", t)
	}
	return true, nil
}

func (e *exprCapture) addExpression(v expression, p *parser) (bool, error) {
	s := "Capture doesn't accept subexpressions (%s)"
	p.logger.Panic(s, v)

	// this should not be reached, but well... fail.
	return false, log.NewError(s, v)
}
