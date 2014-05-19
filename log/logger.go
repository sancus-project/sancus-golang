package log

import (
	"fmt"
	"os"
)

// Logger
type Logger struct {
	Level LogLevel
	Tag   string
}

func (l *Logger) IsLoggable(level LogLevel) bool {
	return (l.Level <= level)
}

func (l *Logger) Printf(level LogLevel, format string, a ...interface{}) (int, error) {
	if l.Level <= level {
		s := fmt.Sprintf(format, a)
		if s != "" {
			return l.Write(level, l.Tag, s)
		}
	}
	return 0, nil
}

// Backend
func (l *Logger) Write(level LogLevel, tag string, message string) (int, error) {
	s := fmt.Sprintf("%s: %s\n", tag, message)
	return os.Stderr.WriteString(s)
}

// Shortcuts
func (l *Logger) Verbose(format string, a ...interface{}) (int, error) {
	return l.Printf(VERBOSE, format, a)
}
func (l *Logger) Debug(format string, a ...interface{}) (int, error) {
	return l.Printf(DEBUG, format, a)
}
func (l *Logger) Info(format string, a ...interface{}) (int, error) {
	return l.Printf(INFO, format, a)
}
func (l *Logger) Warn(format string, a ...interface{}) (int, error) {
	return l.Printf(WARN, format, a)
}
func (l *Logger) Error(format string, a ...interface{}) (int, error) {
	return l.Printf(ERROR, format, a)
}
func (l *Logger) WTF(format string, a ...interface{}) (int, error) {
	return l.Printf(WTF, format, a)
}
func (l *Logger) Fatal(format string, a ...interface{}) {
	l.Printf(ASSERT, format, a)
	os.Exit(1)
}
func (l *Logger) Panic(format string, a ...interface{}) {
	l.Printf(ASSERT, format, a)
	format = fmt.Sprintf("%s: %s\n", l.Tag, format)
	panic(fmt.Sprintf(format, a))
}
