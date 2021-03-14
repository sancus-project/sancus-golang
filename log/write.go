package log

import (
	"bufio"
)

// WriteLines writes log lines
func (logger *Logger) WriteLines(v Variant, lines []string) error {
	return logger.ctx.WriteLines(v, lines, logger.Flags())
}

// WriteLines writes log lines
func (ctx *LoggerContext) WriteLines(v Variant, lines []string, flags uint) error {
	if !ctx.VariantEnabled(v, flags) {
		// Variant disabled
		return nil
	}

	ctx.mu.Lock()
	defer ctx.mu.Unlock()

	w := bufio.NewWriter(ctx.Backend())
	if flags == 0 {
		flags = ctx.Flags()
	}
	for _, s := range lines {
		writeTimestamp(w, flags, ctx.TimeContext())
		w.WriteString(s)
		w.WriteRune('\n')
	}

	return w.Flush()
}
