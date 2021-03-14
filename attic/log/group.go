package log

import (
	"fmt"
	"sync"
)

type Group struct {
	sync.Mutex

	Level   LogLevel
	Backend LoggerBackend
	m       map[string]*Logger
}

// Constructor
func NewGroup(level LogLevel, backend LoggerBackend) *Group {
	return &Group{
		Backend: backend,
		Level:   level,
		m:       make(map[string]*Logger),
	}
}

// Methods
func (m *Group) Get(tag string, a ...interface{}) *Logger {
	return m.Get2(m.Level, tag, a...)
}

func (m *Group) Get2(level LogLevel, tag string, a ...interface{}) *Logger {
	if len(a) > 0 {
		tag = fmt.Sprintf(tag, a...)
	}

	m.Lock()
	defer m.Unlock()

	logger, ok := m.m[tag]
	if !ok {
		logger = NewLogger(tag, level, m)
		m.m[tag] = logger
	}
	return logger
}
