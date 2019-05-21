package log

// io.Writer
func (logger *Logger) Write(data []byte) (int, error) {
	depth := 0 // disable call stack info
	prefix := ""

	if err := logger.WriteLines(logger.FormatBytes(depth, prefix, data)); err != nil {
		return 0, err
	} else {
		return len(data), err
	}
}
