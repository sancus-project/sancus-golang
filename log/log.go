package log

import (
	"fmt"
)

// LogLevel
type LogLevel int

const (
	DEBUG LogLevel = iota + 2
	TRACE
	VERBOSE
	INFO
	WARN
	ERROR
	WTF
	ASSERT
)

var loggers = NewGroup(INFO)

func SetLevel(l LogLevel) {
	loggers.Level = l
}
func Level() LogLevel {
	return loggers.Level
}
func GetLogger(tag string, a ...interface{}) *Logger {
	if len(a) > 0 {
		return loggers.Get(fmt.Sprintf(tag, a...))
	}
	return loggers.Get(tag)
}

// Shortcuts
func Info(tag string, fmt string, a ...interface{}) bool {
	StderrBackend.LogWrite(INFO, tag, fmt, a...)
	return true
}
func Warn(tag string, fmt string, a ...interface{}) bool {
	StderrBackend.LogWrite(WARN, tag, fmt, a...)
	return true
}
func Error(tag string, fmt string, a ...interface{}) bool {
	StderrBackend.LogWrite(ERROR, tag, fmt, a...)
	return true
}
