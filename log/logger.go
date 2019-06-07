package log // import "go.sancus.dev/sancus/log"

import (
	"github.com/kr/pretty"

	"fmt"
)

type Logger struct {
	prefix string
	flags  uint
	ctx    *LoggerContext
}

// fmt.Sprint
func (l *Logger) Print(args ...interface{}) error {
	v := l.DefaultVariant()
	return l.Output(1, v, args...)
}

func (l *Logger) Error(args ...interface{}) error {
	v := l.ErrorVariant()
	return l.Output(1, v, args...)
}

func (l *Logger) Output(calldepth int, v Variant, args ...interface{}) error {
	var s string
	if len(args) > 0 {
		s = fmt.Sprint(args...)
	}
	return l.WriteLines(v, l.Format(deeper(calldepth), v, "", s))
}

func (l *Logger) Output2(calldepth int, v Variant, p string, args ...interface{}) error {
	var s string
	if len(args) > 0 {
		s = fmt.Sprint(args...)
	}
	return l.WriteLines(v, l.Format(deeper(calldepth), v, p, s))
}

// pretty.Sprint
func (l *Logger) PrettyPrint(args ...interface{}) error {
	v := l.DefaultVariant()
	return l.OutputPretty(1, v, args...)
}

func (l *Logger) OutputPretty(calldepth int, v Variant, args ...interface{}) error {
	var s string
	if len(args) > 0 {
		s = pretty.Sprint(args...)
	}
	return l.WriteLines(v, l.Format(deeper(calldepth), v, "", s))
}

func (l *Logger) OutputPretty2(calldepth int, v Variant, p string, args ...interface{}) error {
	var s string
	if len(args) > 0 {
		s = pretty.Sprint(args...)
	}
	return l.WriteLines(v, l.Format(deeper(calldepth), v, p, s))
}

// fmt.Sprintf
func (l *Logger) Printf(fmt string, args ...interface{}) error {
	v := l.DefaultVariant()
	return l.Outputf(1, v, fmt, args...)
}

func (l *Logger) Errorf(fmt string, args ...interface{}) error {
	v := l.ErrorVariant()
	return l.Outputf(1, v, fmt, args...)
}

func (l *Logger) Outputf(calldepth int, v Variant, s string, args ...interface{}) error {
	if len(args) > 0 {
		s = fmt.Sprintf(s, args...)
	}
	return l.WriteLines(v, l.Format(deeper(calldepth), v, "", s))
}

func (l *Logger) Outputf2(calldepth int, v Variant, p string, s string, args ...interface{}) error {
	if len(args) > 0 {
		s = fmt.Sprintf(s, args...)
	}
	return l.WriteLines(v, l.Format(deeper(calldepth), v, p, s))
}

// pretty.Sprintf
func (l *Logger) PrettyPrintf(fmt string, args ...interface{}) error {
	v := l.DefaultVariant()
	return l.OutputPrettyf(1, v, fmt, args...)
}

func (l *Logger) OutputPrettyf(calldepth int, v Variant, s string, args ...interface{}) error {
	if len(args) > 0 {
		s = pretty.Sprintf(s, args...)
	}
	return l.WriteLines(v, l.Format(deeper(calldepth), v, "", s))
}

func (l *Logger) OutputPrettyf2(calldepth int, v Variant, p string, s string, args ...interface{}) error {
	if len(args) > 0 {
		s = pretty.Sprintf(s, args...)
	}
	return l.WriteLines(v, l.Format(deeper(calldepth), v, p, s))
}

// Fatal logs an error and then panics
func (l *Logger) Fatal(err error) {
	if err != nil {
		v := l.ErrorVariant()
		l.Output(1, v, err)
		panic(err)
	}
}
