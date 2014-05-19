package log

import (
	"sync"
)

type LoggerMap struct {
	sync.Mutex

	Level LogLevel
	m     map[string]*Logger
}

// Constructor
func NewLoggerMap(level LogLevel) *LoggerMap {
	return &LoggerMap{Level: level, m: make(map[string]*Logger)}
}

// Methods
func (m *LoggerMap) Get(tag string) *Logger {
	m.Lock()
	logger, ok := m.m[tag]
	if !ok {
		logger = NewLogger(tag, m.Level)
		m.m[tag] = logger
	}
	m.Unlock()
	return logger
}
