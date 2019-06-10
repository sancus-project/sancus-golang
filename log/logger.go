package log // import "go.sancus.dev/sancus/log"

import (
	"github.com/kr/pretty"

	"fmt"
	"os"
)

type Logger struct {
	prefix string
	flags  uint
	ctx    *LoggerContext
}

// Print is equivalent to fmt.Sprint() using the DefaultVariant of the logger
func (l *Logger) Print(args ...interface{}) error {
	return l.Output2(1, l.DefaultVariant(), "", args...)
}

// Error is equivalent to fmt.Sprint() using the ErrorVariant of the logger
func (l *Logger) Error(args ...interface{}) error {
	return l.Output2(1, l.ErrorVariant(), "", args...)
}

// Fatal is equivalent to Error() followed by a call to os.Exit(1)
func (l *Logger) Fatal(args ...interface{}) {
	l.OutputFatal2(1, l.ErrorVariant(), "", args...)
}

// Output is equivalent to fmt.Sprint() when the given variant is enabled
func (l *Logger) Output(calldepth int, v Variant, args ...interface{}) error {
	return l.Output2(deeper(calldepth), v, "", args...)
}

// Output2 is equivalent to fmt.Sprint() with a given prefix when the given variant is enabled
func (l *Logger) Output2(calldepth int, v Variant, p string, args ...interface{}) error {
	var s string

	if !l.VariantEnabled(v) {
		return nil
	}

	if len(args) > 0 {
		s = fmt.Sprint(args...)
	}

	return l.WriteLines(v, l.Format(deeper(calldepth), v, p, s))
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

// pretty.Sprint
func (l *Logger) PrettyPrint(args ...interface{}) error {
	return l.OutputPretty2(1, l.DefaultVariant(), "", args...)
}

func (l *Logger) OutputPretty(calldepth int, v Variant, args ...interface{}) error {
	return l.OutputPretty2(deeper(calldepth), v, "", args...)
}

func (l *Logger) OutputPretty2(calldepth int, v Variant, p string, args ...interface{}) error {
	var s string

	if !l.VariantEnabled(v) {
		return nil
	}

	if len(args) > 0 {
		s = pretty.Sprint(args...)
	}

	return l.WriteLines(v, l.Format(deeper(calldepth), v, p, s))
}

// fmt.Sprintf
func (l *Logger) Printf(fmt string, args ...interface{}) error {
	return l.Outputf2(1, l.DefaultVariant(), "", fmt, args...)
}

func (l *Logger) Errorf(fmt string, args ...interface{}) error {
	return l.Outputf2(1, l.ErrorVariant(), "", fmt, args...)
}

// Fatal is equivalent to Errorf() followed by a call to os.Exit(1)
func (l *Logger) Fatalf(fmt string, args ...interface{}) {
	l.OutputFatalf2(1, l.ErrorVariant(), "", fmt, args...)
}

func (l *Logger) Outputf(calldepth int, v Variant, s string, args ...interface{}) error {
	return l.Outputf2(deeper(calldepth), v, "", s, args...)
}

func (l *Logger) Outputf2(calldepth int, v Variant, p string, s string, args ...interface{}) error {
	if !l.VariantEnabled(v) {
		return nil
	}

	if len(args) > 0 {
		s = fmt.Sprintf(s, args...)
	}

	return l.WriteLines(v, l.Format(deeper(calldepth), v, p, s))
}

// OutputFatalf2 is an unconditional equivalent to Outputf2() followed by a call to os.Exit(1)
func (l *Logger) OutputFatalf2(calldepth int, v Variant, p string, s string, args ...interface{}) {
	if len(args) > 0 {
		s = fmt.Sprintf(s, args...)
	}

	l.WriteLines(v, l.Format(deeper(calldepth), v, p, s))
	os.Exit(1)
}

// pretty.Sprintf
func (l *Logger) PrettyPrintf(fmt string, args ...interface{}) error {
	v := l.DefaultVariant()
	return l.OutputPrettyf(1, v, fmt, args...)
}

func (l *Logger) OutputPrettyf(calldepth int, v Variant, s string, args ...interface{}) error {
	if !l.VariantEnabled(v) {
		return nil
	}

	if len(args) > 0 {
		s = pretty.Sprintf(s, args...)
	}

	return l.WriteLines(v, l.Format(deeper(calldepth), v, "", s))
}

func (l *Logger) OutputPrettyf2(calldepth int, v Variant, p string, s string, args ...interface{}) error {
	if !l.VariantEnabled(v) {
		return nil
	}

	if len(args) > 0 {
		s = pretty.Sprintf(s, args...)
	}

	return l.WriteLines(v, l.Format(deeper(calldepth), v, p, s))
}
