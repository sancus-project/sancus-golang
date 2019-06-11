package log

// Print is equivalent to fmt.Sprint() using the DefaultVariant of the logger
func (l *Logger) Print(args ...interface{}) error {
	return l.Output2(1, l.DefaultVariant(), "", args...)
}

// Error is equivalent to fmt.Sprint() using the ErrorVariant of the logger
func (l *Logger) Error(args ...interface{}) error {
	return l.Output2(1, l.ErrorVariant(), "", args...)
}

// Fatal is equivalent to Error() followed by a call to os.Exit(1)
func (l *Logger) Fatal(args ...interface{}) {
	l.OutputFatal2(1, l.ErrorVariant(), "", args...)
}

// Panic is equivalent to Error() followed by a call to panic()
func (l *Logger) Panic(args ...interface{}) {
	l.OutputPanic2(1, l.ErrorVariant(), "", args...)
}

// Output is equivalent to fmt.Sprint() when the given variant is enabled
func (l *Logger) Output(calldepth int, v Variant, args ...interface{}) error {
	return l.Output2(deeper(calldepth), v, "", args...)
}

// Println() is equivalent to fmt.Sprintln() using the DefaultVariant of the logger
func (l *Logger) Println(args ...interface{}) error {
	return l.Outputln2(1, l.DefaultVariant(), "", args...)
}

// Errorln() is equivalent to fmt.Sprintln() using the ErrorVariant of the logger
func (l *Logger) Errorln(args ...interface{}) error {
	return l.Outputln2(1, l.ErrorVariant(), "", args...)
}

// Fatalln is equivalent to Errorln() followed by a call to os.Exit(1)
func (l *Logger) Fatalln(args ...interface{}) {
	l.OutputFatalln2(1, l.ErrorVariant(), "", args...)
}

// Panicln is equivalent to Errorln() followed by a call to panic()
func (l *Logger) Panicln(args ...interface{}) {
	l.OutputPanicln2(1, l.ErrorVariant(), "", args...)
}

// Outputln is equivalent to fmt.Sprintln() when the given variant is enabled
func (l *Logger) Outputln(calldepth int, v Variant, args ...interface{}) error {
	return l.Outputln2(deeper(calldepth), v, "", args...)
}

// pretty.Sprint
func (l *Logger) PrettyPrint(args ...interface{}) error {
	return l.OutputPretty2(1, l.DefaultVariant(), "", args...)
}

func (l *Logger) OutputPretty(calldepth int, v Variant, args ...interface{}) error {
	return l.OutputPretty2(deeper(calldepth), v, "", args...)
}

// fmt.Sprintf
func (l *Logger) Printf(fmt string, args ...interface{}) error {
	return l.Outputf2(1, l.DefaultVariant(), "", fmt, args...)
}

func (l *Logger) Errorf(fmt string, args ...interface{}) error {
	return l.Outputf2(1, l.ErrorVariant(), "", fmt, args...)
}

// Fatalf is equivalent to Errorf() followed by a call to os.Exit(1)
func (l *Logger) Fatalf(fmt string, args ...interface{}) {
	l.OutputFatalf2(1, l.ErrorVariant(), "", fmt, args...)
}

// Panicf is equivalent to Errorf() followed by a call to panic()
func (l *Logger) Panicf(fmt string, args ...interface{}) {
	l.OutputPanicf2(1, l.ErrorVariant(), "", fmt, args...)
}

func (l *Logger) Outputf(calldepth int, v Variant, s string, args ...interface{}) error {
	return l.Outputf2(deeper(calldepth), v, "", s, args...)
}

// pretty.Sprintln
func (l *Logger) PrettyPrintln(args ...interface{}) error {
	return l.OutputPrettyln2(1, l.DefaultVariant(), "", args...)
}

// pretty.Sprintf
func (l *Logger) PrettyPrintf(fmt string, args ...interface{}) error {
	return l.OutputPrettyf2(1, l.DefaultVariant(), "", fmt, args...)
}

func (l *Logger) OutputPrettyf(calldepth int, v Variant, s string, args ...interface{}) error {
	return l.OutputPrettyf2(deeper(calldepth), v, "", s, args...)
}
