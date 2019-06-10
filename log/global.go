package log

var defaultContext = NewLoggerContext(Lstdflags)
var defaultLogger = defaultContext.NewLogger("")

// New creates a new logger from the global context
func New(prefix string, flags uint) *Logger {
	return defaultContext.NewLogger(prefix).SetFlags(flags)
}

// Print is equivalent to fmt.Sprint() using the DefaultVariant of the logger
func Print(args ...interface{}) {
	defaultLogger.Output2(1, defaultLogger.DefaultVariant(), "", args...)
}

// Error is equivalent to fmt.Sprint() using the ErrorVariant of the logger
func Error(args ...interface{}) {
	defaultLogger.Output2(1, defaultLogger.ErrorVariant(), "", args...)
}

// Fatal is equivalent to Error() followed by a call to os.Exit(1)
func Fatal(args ...interface{}) {
	defaultLogger.OutputFatal2(1, defaultLogger.ErrorVariant(), "", args...)
}

// Panic is equivalent to Error() followed by a call to panic()
func Panic(args ...interface{}) {
	defaultLogger.OutputPanic2(1, defaultLogger.ErrorVariant(), "", args...)
}

// Println is equivalent to fmt.Sprintln() using the DefaultVariant of the logger
func Println(args ...interface{}) {
	defaultLogger.Outputln2(1, defaultLogger.DefaultVariant(), "", args...)
}

// Errorln is equivalent to fmt.Sprintln() using the ErrorVariant of the logger
func Errorln(args ...interface{}) {
	defaultLogger.Outputln2(1, defaultLogger.ErrorVariant(), "", args...)
}

// Fatalln is equivalent to Error() followed by a call to os.Exit(1)
func Fatalln(args ...interface{}) {
	defaultLogger.OutputFatalln2(1, defaultLogger.ErrorVariant(), "", args...)
}

// Panicln is equivalent to Error() followed by a call to panic()
func Panicln(args ...interface{}) {
	defaultLogger.OutputPanicln2(1, defaultLogger.ErrorVariant(), "", args...)
}

// Printf is equivalent to fmt.Sprintf() using the DefaultVariant of the logger
func Printf(s string, args ...interface{}) {
	defaultLogger.Outputf2(1, defaultLogger.DefaultVariant(), "", s, args...)
}

// Errorf is equivalent to fmt.Sprintf() using the ErrorVariant of the logger
func Errorf(s string, args ...interface{}) {
	defaultLogger.Outputf2(1, defaultLogger.ErrorVariant(), "", s, args...)
}

// Fatalf is equivalent to Errorf() followed by a call to os.Exit(1)
func Fatalf(s string, args ...interface{}) {
	defaultLogger.OutputFatalf2(1, defaultLogger.ErrorVariant(), "", s, args...)
}

// Panicf is equivalent to Errorf() followed by a call to panic()
func Panicf(s string, args ...interface{}) {
	defaultLogger.OutputPanicf2(1, defaultLogger.ErrorVariant(), "", s, args...)
}
