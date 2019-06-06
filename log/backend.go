package log

import (
	"io"
	"os"
)

// SetDefaultBackend
var defaultBackend io.Writer = os.Stderr

func SetDefaultBackend(w io.Writer) {
	if w == nil {
		defaultBackend = os.Stderr
	} else {
		defaultBackend = w
	}
}

// SetBackend
func (ctx *LoggerContext) SetBackend(w io.Writer) *LoggerContext {
	ctx.backend = w
	return ctx
}

func (logger *Logger) SetBackend(w io.Writer) *Logger {
	logger.ctx.SetBackend(w)
	return logger
}

// Backend
func (ctx *LoggerContext) Backend() io.Writer {
	if ctx.backend == nil {
		return defaultBackend
	} else {
		return ctx.backend
	}
}

func (logger *Logger) Backend() io.Writer {
	return logger.ctx.Backend()
}
