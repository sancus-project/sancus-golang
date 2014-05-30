package log

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

var loggers = NewGroup(INFO, &StderrBackend)

func SetLevel(l LogLevel) {
	loggers.Level = l
}
func Level() LogLevel {
	return loggers.Level
}
func GetLogger(tag string, a ...interface{}) *Logger {
	return loggers.Get(tag, a...)
}

// Shortcuts
func Info(tag string, fmt string, a ...interface{}) bool {
	loggers.Backend.LogWrite(INFO, tag, fmt, a...)
	return true
}
func Warn(tag string, fmt string, a ...interface{}) bool {
	loggers.Backend.LogWrite(WARN, tag, fmt, a...)
	return true
}
func Error(tag string, fmt string, a ...interface{}) bool {
	loggers.Backend.LogWrite(ERROR, tag, fmt, a...)
	return true
}
