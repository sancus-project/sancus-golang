package log

// io.Writer
func (logger *Logger) Write(data []byte) (int, error) {
	v := logger.DefaultVariant()

	if !(logger.VariantEnabled(v)) {
		// DefaultVariant disabled
		return 0, nil
	}

	depth := 0 // disable call stack info
	prefix := ""

	if err := logger.WriteLines(v, logger.FormatBytes(depth, v, prefix, data)); err != nil {
		return 0, err
	} else {
		return len(data), err
	}
}
