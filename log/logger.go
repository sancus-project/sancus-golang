package log

import (
	"os"
)

// Logger
type Logger struct {
	group *Group
	Level LogLevel
	tag   string
}

func NewLogger(tag string, level LogLevel, group *Group) *Logger {
	return &Logger{
		tag:   tag,
		Level: level,
		group: group,
	}
}

func (l *Logger) IsLoggable(level LogLevel) bool {
	return (l.Level <= level)
}

func (l *Logger) Tag() string {
	return l.tag
}

func (l *Logger) Printf(level LogLevel, fmt string, a ...interface{}) bool {
	if l.IsLoggable(level) {
		l.group.Backend.LogWrite(level, l.tag, fmt, a...)
		return true
	}
	return false
}

// Shortcuts
func (l *Logger) Debug(format string, a ...interface{}) bool {
	return l.Printf(DEBUG, format, a...)
}
func (l *Logger) Trace(format string, a ...interface{}) bool {
	return l.Printf(TRACE, format, a...)
}
func (l *Logger) Verbose(format string, a ...interface{}) bool {
	return l.Printf(VERBOSE, format, a...)
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
	panic(StderrBackend.Encode(ASSERT, l.tag, format, a...))
}
