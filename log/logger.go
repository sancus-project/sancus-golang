package log

import (
	"fmt"
	"os"
)

func StderrLogWrite(level LogLevel, tag string, fmt string, a ...interface{}) (int, error) {
	return os.Stderr.WriteString(BaseEncoder(level, tag, fmt, a...))
}

func BaseEncoder(level LogLevel, tag string, format string, a ...interface{}) string {
	var s string

	switch level {
	case VERBOSE:
		s = "V/"
	case DEBUG:
		s = "D/"
	case INFO:
		s = "I/"
	case WARN:
		s = "W/"
	case ERROR:
		s = "E/"
	case WTF:
		s = "F/"
	case ASSERT:
		s = "A/"
	default:
		s = ""
	}

	if s == "" && tag == "" {
		// NOP
	} else if tag != "" {
		s += tag + ": "
	} else {
		s += "undefined: "
	}

	if len(a) > 0 {
		s += fmt.Sprintf(format, a...)
	} else {
		s += format
	}

	return s + "\n"
}

// Logger
type Logger struct {
	Level LogLevel
	tag   string
}

func NewLogger(tag string, level LogLevel) *Logger {
	return &Logger{tag: tag, Level: level}
}

func (l *Logger) IsLoggable(level LogLevel) bool {
	return (l.Level <= level)
}

func (l *Logger) Tag() string {
	return l.tag
}

func (l *Logger) Printf(level LogLevel, fmt string, a ...interface{}) (int, error) {
	if l.IsLoggable(level) {
		return StderrLogWrite(level, l.tag, fmt, a...)
	}
	return 0, nil
}

// Shortcuts
func (l *Logger) Verbose(format string, a ...interface{}) (int, error) {
	return l.Printf(VERBOSE, format, a...)
}
func (l *Logger) Debug(format string, a ...interface{}) (int, error) {
	return l.Printf(DEBUG, format, a...)
}
func (l *Logger) Info(format string, a ...interface{}) (int, error) {
	return l.Printf(INFO, format, a...)
}
func (l *Logger) Warn(format string, a ...interface{}) (int, error) {
	return l.Printf(WARN, format, a...)
}
func (l *Logger) Error(format string, a ...interface{}) (int, error) {
	return l.Printf(ERROR, format, a...)
}
func (l *Logger) WTF(format string, a ...interface{}) (int, error) {
	return l.Printf(WTF, format, a...)
}
func (l *Logger) Fatal(format string, a ...interface{}) {
	l.Printf(ASSERT, format, a...)
	os.Exit(1)
}
func (l *Logger) Panic(format string, a ...interface{}) {
	l.Printf(ASSERT, format, a...)
	panic(BaseEncoder(ASSERT, l.tag, format, a...))
}
