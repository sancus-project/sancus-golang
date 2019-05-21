package log

import (
	"bufio"
	"io"
	"os"
)

// backend
var output io.Writer = os.Stderr

// SetOutput sets backend writer
func (logger *Logger) SetOutput(w io.Writer) *Logger {
	if w == nil {
		output = os.Stderr
	} else {
		output = w
	}

	return logger
}

// write log lines
func (logger *Logger) WriteLines(lines []string) error {
	logger.mu.Lock()
	defer logger.mu.Unlock()

	w := bufio.NewWriter(output)
	for _, s := range lines {
		w.WriteString(s)
		w.WriteRune('\n')
	}

	return w.Flush()
}
