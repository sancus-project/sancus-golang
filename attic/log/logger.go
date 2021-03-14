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

func (l *Logger) SubLogger(tag string, a ...interface{}) *Logger {
	if len(a) > 0 {
		tag = "%s" + tag
		a = append([]interface{}{l.tag}, a...)
	} else {
		tag = l.tag + tag
	}

	return l.group.Get2(l.Level, tag, a...)
}

func (l *Logger) IsLoggable(level LogLevel) bool {
	return (l.Level <= level)
}

func (l *Logger) Tag() string {
	return l.tag
}

func (l *Logger) Log(level LogLevel, fmt string, a ...interface{}) bool {
	if l.IsLoggable(level) {
		l.group.Backend.LogWrite(level, l.tag, fmt, a...)
		return true
	}
	return false
}

// Shortcuts
func (l *Logger) Debug(format string, a ...interface{}) bool {
	return l.Log(DEBUG, format, a...)
}
func (l *Logger) Trace(format string, a ...interface{}) bool {
	return l.Log(TRACE, format, a...)
}
func (l *Logger) Verbose(format string, a ...interface{}) bool {
	return l.Log(VERBOSE, format, a...)
}
func (l *Logger) Info(format string, a ...interface{}) bool {
	return l.Log(INFO, format, a...)
}
func (l *Logger) Warn(format string, a ...interface{}) bool {
	return l.Log(WARN, format, a...)
}
func (l *Logger) Error(format string, a ...interface{}) bool {
	return l.Log(ERROR, format, a...)
}
func (l *Logger) WTF(format string, a ...interface{}) bool {
	return l.Log(WTF, format, a...)
}
func (l *Logger) Fatal(format string, a ...interface{}) {
	l.Log(ASSERT, format, a...)
	os.Exit(1)
}
func (l *Logger) Panic(format string, a ...interface{}) {
	l.Log(ASSERT, format, a...)
	panic(StderrBackend.Encode(ASSERT, l.tag, format, a...))
}
