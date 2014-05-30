package log

import (
	"fmt"
	"sync"
)

type Group struct {
	sync.Mutex

	Level LogLevel
	m     map[string]*Logger
}

// Constructor
func NewGroup(level LogLevel) *Group {
	return &Group{Level: level, m: make(map[string]*Logger)}
}

// Methods
func (m *Group) Get(tag string, a ...interface{}) *Logger {
	if len(a) > 0 {
		tag = fmt.Sprintf(tag, a...)
	}

	m.Lock()
	logger, ok := m.m[tag]
	if !ok {
		logger = NewLogger(tag, m.Level)
		m.m[tag] = logger
	}
	m.Unlock()
	return logger
}
