package log

// LogLevel
type LogLevel int

const (
	VERBOSE LogLevel = 2
	DEBUG
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
