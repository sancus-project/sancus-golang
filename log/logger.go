package log

import (
	"bytes"
	"fmt"
	"os"
)

func StderrLogWrite(level LogLevel, tag string, fmt string, a ...interface{}) (int, error) {
	return os.Stderr.Write(BaseEncoder(level, tag, fmt, a...))
}

/* Returns line representation of a log entry */
func BaseEncoder(level LogLevel, tag string, format string, a ...interface{}) []byte {
	var b bytes.Buffer
	var c rune

	switch level {
	case VERBOSE:
		c = 'V'
	case DEBUG:
		c = 'D'
	case INFO:
		c = 'I'
	case WARN:
		c = 'W'
	case ERROR:
		c = 'E'
	case WTF:
		c = 'F'
	case ASSERT:
		c = 'A'
	default:
		c = '\000'
	}

	if tag != "" {
		if c != '\000' {
			b.WriteRune(c)
			b.WriteRune('/')
		}
		b.WriteString(tag)
		b.WriteString(": ")
	} else if c != '\000' {
		b.WriteRune(c)
		b.WriteString(": ")
	}

	if len(a) > 0 {
		b.WriteString(fmt.Sprintf(format, a...))
	} else {
		b.WriteString(format)
	}

	b.WriteRune('\n')
	return b.Bytes()
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

func (l *Logger) Printf(level LogLevel, fmt string, a ...interface{}) bool {
	if l.IsLoggable(level) {
		StderrLogWrite(level, l.tag, fmt, a...)
		return true
	}
	return false
}

// Shortcuts
func (l *Logger) Verbose(format string, a ...interface{}) bool {
	return l.Printf(VERBOSE, format, a...)
}
func (l *Logger) Debug(format string, a ...interface{}) bool {
	return l.Printf(DEBUG, format, a...)
}
func (l *Logger) Info(format string, a ...interface{}) bool {
	return l.Printf(INFO, format, a...)
}
func (l *Logger) Warn(format string, a ...interface{}) bool {
	return l.Printf(WARN, format, a...)
}
func (l *Logger) Error(format string, a ...interface{}) bool {
	return l.Printf(ERROR, format, a...)
}
func (l *Logger) WTF(format string, a ...interface{}) bool {
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
