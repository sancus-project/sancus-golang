package log // import "go.sancus.dev/sancus/log"

import (
	"github.com/kr/pretty"

	"fmt"
	"os"
	"strings"
)

type Logger struct {
	prefix string
	flags  uint
	ctx    *LoggerContext
}

// Output2 is equivalent to fmt.Sprint() with a given prefix when the given variant is enabled
func (l *Logger) Output2(calldepth int, v Variant, p string, args ...interface{}) error {
	var err error

	if l.VariantEnabled(v) {
		var s string

		if len(args) > 0 {
			s = fmt.Sprint(args...)
		}

		err = l.WriteLines(v, l.Format(deeper(calldepth), v, p, s))
	}

	return err
}

// OutputFatal2 is an unconditional equivalent to Output2() followed by a call to os.Exit(1)
func (l *Logger) OutputFatal2(calldepth int, v Variant, p string, args ...interface{}) {
	var s string

	if len(args) > 0 {
		s = fmt.Sprint(args...)
	}

	l.WriteLines(v, l.Format(deeper(calldepth), v, p, s))
	os.Exit(1)
}

// OutputPanic2 is an unconditional equivalent to Output2() followed by a call to panic()
func (l *Logger) OutputPanic2(calldepth int, v Variant, p string, args ...interface{}) {
	var s string

	if len(args) > 0 {
		s = fmt.Sprint(args...)
	}

	lines := l.Format(deeper(calldepth), v, p, s)
	l.WriteLines(v, lines)

	panic(strings.Join(lines, "\n"))
}

// Outputln2 is equivalent to fmt.Sprint() with a given prefix when the given variant is enabled
func (l *Logger) Outputln2(calldepth int, v Variant, p string, args ...interface{}) error {
	var err error

	if l.VariantEnabled(v) {
		var s string

		if len(args) > 0 {
			s = fmt.Sprintln(args...)
		}

		err = l.WriteLines(v, l.Format(deeper(calldepth), v, p, s))
	}

	return err
}

// OutputFatalln2() is an unconditional equivalent to Outputln2() followed by a call to os.Exit(1)
func (l *Logger) OutputFatalln2(calldepth int, v Variant, p string, args ...interface{}) {
	var s string

	if len(args) > 0 {
		s = fmt.Sprintln(args...)
	}

	l.WriteLines(v, l.Format(deeper(calldepth), v, p, s))
	os.Exit(1)
}

// OutputPanicln2() is an unconditional equivalent to Outputln2() followed by a call to panic()
func (l *Logger) OutputPanicln2(calldepth int, v Variant, p string, args ...interface{}) {
	var s string

	if len(args) > 0 {
		s = fmt.Sprintln(args...)
	}

	lines := l.Format(deeper(calldepth), v, p, s)
	l.WriteLines(v, lines)

	panic(strings.Join(lines, "\n"))
}

func (l *Logger) OutputPretty2(calldepth int, v Variant, p string, args ...interface{}) error {
	var err error

	if l.VariantEnabled(v) {
		var s string

		if len(args) > 0 {
			s = pretty.Sprint(args...)
		}

		err = l.WriteLines(v, l.Format(deeper(calldepth), v, p, s))
	}

	return err
}

func (l *Logger) Outputf2(calldepth int, v Variant, p string, s string, args ...interface{}) error {
	var err error

	if l.VariantEnabled(v) {
		if len(args) > 0 {
			s = fmt.Sprintf(s, args...)
		}

		err = l.WriteLines(v, l.Format(deeper(calldepth), v, p, s))
	}

	return err
}

// OutputFatalf2 is an unconditional equivalent to Outputf2() followed by a call to os.Exit(1)
func (l *Logger) OutputFatalf2(calldepth int, v Variant, p string, s string, args ...interface{}) {
	if len(args) > 0 {
		s = fmt.Sprintf(s, args...)
	}

	l.WriteLines(v, l.Format(deeper(calldepth), v, p, s))
	os.Exit(1)
}

// OutputPanicf2 is an unconditional equivalent to Outputf2() followed by a call to panic()
func (l *Logger) OutputPanicf2(calldepth int, v Variant, p string, s string, args ...interface{}) {
	if len(args) > 0 {
		s = fmt.Sprintf(s, args...)
	}

	lines := l.Format(deeper(calldepth), v, p, s)
	l.WriteLines(v, lines)

	panic(strings.Join(lines, "\n"))
}

func (l *Logger) OutputPrettyf2(calldepth int, v Variant, p string, s string, args ...interface{}) error {
	var err error

	if l.VariantEnabled(v) {
		if len(args) > 0 {
			s = pretty.Sprintf(s, args...)
		}
		err = l.WriteLines(v, l.Format(deeper(calldepth), v, p, s))
	}

	return err
}
