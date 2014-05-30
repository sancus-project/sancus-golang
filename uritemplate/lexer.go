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
	tokenColon        // ":"
	tokenPipe         // "|"
	tokenIdentifier   // [a-zA-Z] [a-zA-Z0-9_]*
	tokenText         // [a-zA-Z0-9.,%_-]
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

func lexText(l *lexer) stateFn {
	l.log.Trace("state:lexText, %v..%v/%v",
		l.start, l.pos, len(l.input))

	for {
		var r rune
		var t tokenType
		var nextState = lexText

		if r = l.next(); r == eof {
			break
		}

		l.log.Debug("state:lexText r:%c (%v) [%v..%v] ",
			r, l.width,
			l.start, l.pos)

		switch r {
		case '[':
			t = tokenLeftBracket
		case ']':
			t = tokenRightBracket
		case '$':
			t = tokenEOL
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
