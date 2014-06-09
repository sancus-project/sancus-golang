package uritemplate

import (
	"fmt"
	"go.sancus.io/core/log"
	"unicode/utf8"
)

/*
 * Based on "Lexical Scanning in Go" by Rob Pike
 * http://cuddle.googlecode.com/hg/talk/lex.html
 *
 */

/*
 * Optional = "[" ... "]"
 * Capture  = "{" name [ ":" option { "|" option } ] "}"
 * EOL      = "$"
 */

/*
 * Token
 */
type tokenType int

type token struct {
	typ tokenType
	val string
}

const (
	tokenError tokenType = iota
	tokenEOF
	tokenEOL          // "$"
	tokenLeftBracket  // "["
	tokenRightBracket // "]"
	tokenLeftBrace    // "{"
	tokenRightBrace   // "}"
	tokenIdentifier   // [a-zA-Z] [a-zA-Z0-9_]*
	tokenText         // [a-zA-Z0-9.,%_-]
	tokenOption
)

const (
	eof = -1
)

func (t token) String() string {
	switch t.typ {
	case tokenEOF:
		return "EOF"
	case tokenError:
		return "Error:" + t.val
	case tokenText:
		return fmt.Sprintf("LITERAL:%q", t.val)
	case tokenIdentifier:
		return fmt.Sprintf("ID:%s", t.val)
	case tokenOption:
		return fmt.Sprintf("OPTION:%q", t.val)
	default:
		return fmt.Sprintf("%q", t.val)
	}
}

/*
 * Lexer
 */
type lexer struct {
	log    *log.Logger // logger
	input  string      // string being scanned
	start  int         // start position of this token
	pos    int         // current position in the input
	width  int         // width of the last run read from input
	state  stateFn     // current state
	tokens chan token  // channel of scanned tokens
}

func newLexer(input string, logger *log.Logger) *lexer {
	l := &lexer{
		log:    logger,
		input:  input,
		state:  lexText,
		tokens: make(chan token, 2),
	}

	return l
}

func (l *lexer) nextToken() token {
	for {
		select {
		case t := <-l.tokens:
			return t
		default:
			l.state = l.state(l)
		}
	}
	l.log.Panic("not reached")
	return token{}
}

func (l *lexer) emit(typ tokenType) {
	s := l.input[l.start:l.pos]
	t := token{typ, s}

	l.log.Trace("emit(%s) [%v..%v]", t, l.start, l.pos)

	l.tokens <- t
	l.start = l.pos
}

func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}

	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) restore() {
	l.pos += l.width
}

func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

/*
 * States
 */
type stateFn func(*lexer) stateFn

func (l *lexer) fail(msg string, a ...interface{}) stateFn {
	if len(a) > 0 {
		msg = fmt.Sprintf(msg, a...)
	}

	l.tokens <- token{tokenError, msg}
	return nil
}

func lexText(l *lexer) stateFn {
	l.log.Trace("state:lexText [%v..%v/%v]",
		l.start, l.pos, len(l.input))

	for {
		var r rune
		var t tokenType
		var nextState = lexText

		if r = l.next(); r == eof {
			break
		}

		l.log.Debug("state:lexText r:%q (%v) [%v..%v/%v] ",
			r, l.width,
			l.start, l.pos, len(l.input))

		switch r {
		case '[':
			t = tokenLeftBracket
		case ']':
			t = tokenRightBracket
		case '$':
			t = tokenEOL
		case '{':
			t, nextState = tokenLeftBrace, lexCaptureID
		default:
			continue
		}

		l.backup()
		if l.pos > l.start {
			l.emit(tokenText)
		}

		l.restore()
		l.emit(t)

		return nextState
	}
	if l.pos > l.start {
		l.emit(tokenText)
	}
	l.emit(tokenEOF)
	return nil
}

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
