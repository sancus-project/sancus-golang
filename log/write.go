package log

import (
	"bufio"
)

// WriteLines writes log lines
func (logger *Logger) WriteLines(lines []string) error {
	return logger.ctx.WriteLines(lines)
}

// WriteLines writes log lines
func (ctx *LoggerContext) WriteLines(lines []string) error {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()

	w := bufio.NewWriter(ctx.Backend())
	for _, s := range lines {
		w.WriteString(s)
		w.WriteRune('\n')
	}

	return w.Flush()
}
