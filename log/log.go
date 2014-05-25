package log

// LogLevel
type LogLevel int

const (
	DEBUG LogLevel = iota + 2
	VERBOSE
	INFO
	WARN
	ERROR
	WTF
	ASSERT
)

// LoggerMap
var loggers = NewLoggerMap(INFO)

func SetLevel(l LogLevel) {
	loggers.Level = l
}
func Level() LogLevel {
	return loggers.Level
}
func GetLogger(tag string) *Logger {
	return loggers.Get(tag)
}

// Shortcuts
func Info(tag string, fmt string, a ...interface{}) (int, error) {
	return StderrLogWrite(INFO, tag, fmt, a...)
}
func Warn(tag string, fmt string, a ...interface{}) (int, error) {
	return StderrLogWrite(WARN, tag, fmt, a...)
}
func Error(tag string, fmt string, a ...interface{}) (int, error) {
	return StderrLogWrite(ERROR, tag, fmt, a...)
}
