package log

type writer struct {
	prefix string
	logger *Logger
	v      Variant
}

func (logger *Logger) NewWriter(prefix string, v Variant) *writer {
	return &writer{
		prefix: prefix,
		logger: logger,
		v:      v,
	}
}

func (w *writer) Write(data []byte) (l int, err error) {
	if !w.logger.VariantEnabled(w.v) {
		// NOP
	} else if err = w.logger.WriteLines(w.v, w.logger.FormatBytes(0, w.v, w.prefix, data)); err == nil {
		l = len(data)
	}

	return
}

// io.Writer
func (logger *Logger) Write(data []byte) (l int, err error) {
	v := logger.DefaultVariant()

	if !logger.VariantEnabled(v) {
		// NOP
	} else if err = logger.WriteLines(v, logger.FormatBytes(0, v, "", data)); err == nil {
		l = len(data)
	}

	return
}
