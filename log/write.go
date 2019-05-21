package log

import (
	"bufio"
	"io"
	"os"
)

// backend
var output io.Writer = os.Stderr

// SetOutput sets backend writer
func (self *Logger) SetOutput(w io.Writer) *Logger {
	if w == nil {
		output = os.Stderr
	} else {
		output = w
	}

	return self
}

// write log lines
func (self *Logger) WriteLines(lines []string) error {
	self.mu.Lock()
	defer self.mu.Unlock()

	w := bufio.NewWriter(output)
	for _, s := range lines {
		w.WriteString(s)
		w.WriteRune('\n')
	}

	return w.Flush()
}
