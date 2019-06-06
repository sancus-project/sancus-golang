package log

import (
	"io"
	"sync"
)

type LoggerContext struct {
	mu      sync.Mutex
	flags   uint
	backend io.Writer
}

func NewLoggerContext(flags uint) *LoggerContext {
	return &LoggerContext{
		flags: apply_flags(0, flags),
	}
}

func (ctx *LoggerContext) NewLogger(prefix string) *Logger {
	return &Logger{
		prefix: prefix,
		ctx:    ctx,
	}
}

//
//
func (ctx *LoggerContext) SetFlags(flags uint) *LoggerContext {
	ctx.flags = apply_flags(ctx.flags, flags)

	return ctx
}

func (ctx *LoggerContext) Flags() uint {
	return ctx.flags
}

//
//
func (ctx *LoggerContext) Lock() {
	ctx.mu.Lock()
}

func (ctx *LoggerContext) Unlock() {
	ctx.mu.Unlock()
}
