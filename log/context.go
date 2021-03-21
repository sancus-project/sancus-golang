package log

import (
	"fmt"
	"io"
	"sync"
)

type LoggerContext struct {
	mu      sync.Mutex
	flags   uint
	backend io.Writer
	timectx TimeContext

	errorVariant Variant
	variants     map[Variant]loggerVariant
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

func (logger *Logger) New(prefix string, args ...interface{}) *Logger {
	if len(args) > 0 {
		prefix = fmt.Sprintf(prefix, args...)
	}
	p := logger.ctx.NewLogger(logger.prefix + prefix)
	p.flags = logger.flags
	p.defaultVariant = logger.defaultVariant
	return p
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
