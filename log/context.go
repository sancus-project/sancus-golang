package log

import (
	"io"
	"sync"
)

type LoggerContext struct {
	mu      sync.Mutex
	flags   uint
	backend io.Writer
	timectx TimeContext

	defaultVariant Variant
	errorVariant   Variant
	variants       map[Variant]loggerVariant
}

func NewLoggerContext(flags uint) *LoggerContext {
	return &LoggerContext{
		flags:    apply_flags(0, flags),
		variants: make(map[Variant]loggerVariant),
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
func (ctx *LoggerContext) SetTimeContext(tctx TimeContext) *LoggerContext {
	ctx.timectx = tctx
	return ctx
}

func (ctx *LoggerContext) TimeContext() TimeContext {
	if ctx.timectx == nil {
		return StdTimeContext
	} else {
		return ctx.timectx
	}
}

//
//
func (ctx *LoggerContext) Lock() {
	ctx.mu.Lock()
}

func (ctx *LoggerContext) Unlock() {
	ctx.mu.Unlock()
}
