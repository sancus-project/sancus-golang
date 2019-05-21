package log // import "github.com/amery/go-misc/log"

import (
	"io"
	"sync"
)

//
type LoggerContext struct {
	mu     sync.Mutex
	output io.Writer
}

func NewLoggerContext() *LoggerContext {
	return &LoggerContext{}
}

func (ctx *LoggerContext) NewLogger(prefix string) *Logger {
	return &Logger{
		prefix: prefix,
		ctx:    ctx,
	}
}

//
type Logger struct {
	mu     sync.Mutex
	prefix string
	ctx    *LoggerContext
}

var defaultContext = &LoggerContext{}

func New(prefix string) *Logger {
	return defaultContext.NewLogger(prefix)
}
