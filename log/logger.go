package log // import "github.com/amery/go-misc/log"

import (
	"github.com/kr/pretty"

	"fmt"
)

type Logger struct {
	prefix string
	ctx    *LoggerContext
}

// fmt.Sprint
func (l *Logger) Print(args ...interface{}) error {
	return l.Output(1, args...)
}

func (l *Logger) Output(calldepth int, args ...interface{}) error {
	var s string
	if len(args) > 0 {
		s = fmt.Sprint(args...)
	}
	return l.WriteLines(l.Format(deeper(calldepth), "", s))
}

func (l *Logger) Output2(calldepth int, p string, args ...interface{}) error {
	var s string
	if len(args) > 0 {
		s = fmt.Sprint(args...)
	}
	return l.WriteLines(l.Format(deeper(calldepth), p, s))
}

// pretty.Sprint
func (l *Logger) PrettyPrint(args ...interface{}) error {
	return l.OutputPretty(1, args...)
}

func (l *Logger) OutputPretty(calldepth int, args ...interface{}) error {
	var s string
	if len(args) > 0 {
		s = pretty.Sprint(args...)
	}
	return l.WriteLines(l.Format(deeper(calldepth), "", s))
}

func (l *Logger) OutputPretty2(calldepth int, p string, args ...interface{}) error {
	var s string
	if len(args) > 0 {
		s = pretty.Sprint(args...)
	}
	return l.WriteLines(l.Format(deeper(calldepth), p, s))
}

// fmt.Sprintf
func (l *Logger) Printf(fmt string, args ...interface{}) error {
	return l.Outputf(1, fmt, args...)
}

func (l *Logger) Outputf(calldepth int, s string, args ...interface{}) error {
	if len(args) > 0 {
		s = fmt.Sprintf(s, args...)
	}
	return l.WriteLines(l.Format(deeper(calldepth), "", s))
}

func (l *Logger) Outputf2(calldepth int, p string, s string, args ...interface{}) error {
	if len(args) > 0 {
		s = fmt.Sprintf(s, args...)
	}
	return l.WriteLines(l.Format(deeper(calldepth), p, s))
}

// pretty.Sprintf
func (l *Logger) PrettyPrintf(fmt string, args ...interface{}) error {
	return l.OutputPrettyf(1, fmt, args...)
}

func (l *Logger) OutputPrettyf(calldepth int, s string, args ...interface{}) error {
	if len(args) > 0 {
		s = pretty.Sprintf(s, args...)
	}
	return l.WriteLines(l.Format(deeper(calldepth), "", s))
}

func (l *Logger) OutputPrettyf2(calldepth int, p string, s string, args ...interface{}) error {
	if len(args) > 0 {
		s = pretty.Sprintf(s, args...)
	}
	return l.WriteLines(l.Format(deeper(calldepth), p, s))
}
