package log

var defaultContext = NewLoggerContext(Lstdflags)

// New creates a new logger from the global context
func New(prefix string, flags uint) *Logger {
	return defaultContext.NewLogger(prefix).SetFlags(flags)
}
