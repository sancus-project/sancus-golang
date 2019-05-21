package log

// io.Writer
func (self *Logger) Write(data []byte) (int, error) {
	depth := 0 // disable call stack info
	prefix := ""

	if err := self.WriteLines(self.FormatBytes(depth, prefix, data)); err != nil {
		return 0, err
	} else {
		return len(data), err
	}
}
