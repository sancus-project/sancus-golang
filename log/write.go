package log

import (
	"bufio"
	"io"
	"os"
)

// backend
var output io.Writer = os.Stderr

func SetDefaultOutput(w io.Writer) {
	if w == nil {
		output = os.Stderr
	} else {
		output = w
	}
}

func (logger *Logger) SetOutput(w io.Writer) *Logger {
	logger.ctx.SetOutput(w)
	return logger
}

func (ctx *LoggerContext) SetOutput(w io.Writer) *LoggerContext {
	ctx.output = w
	return ctx
}

func (ctx *LoggerContext) GetOutput() io.Writer {
	if ctx.output != nil {
		return ctx.output
	} else {
		return output
	}
}

// WriteLines writes log lines
func (logger *Logger) WriteLines(lines []string) error {
	return logger.ctx.WriteLines(lines)
}

func (ctx *LoggerContext) WriteLines(lines []string) error {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()

	w := bufio.NewWriter(ctx.GetOutput())
	for _, s := range lines {
		w.WriteString(s)
		w.WriteRune('\n')
	}

	return w.Flush()
}
