package context

import (
	"net/http"
	"sync"
)

type Context map[string]interface{}

type ContextMap struct {
	sync.RWMutex

	context map[*http.Request]Context
}

// Create entries
func (m *ContextMap) Set(r *http.Request, k string, v interface{}) {
	var ctx Context

	m.Lock()
	if ctx = m.context[r]; ctx == nil {
		ctx = make(Context)
		m.context[r] = ctx
	}
	m.Unlock()

	ctx[k] = v
}

// Get entries
func (m *ContextMap) GetAll(r *http.Request) (Context, bool) {
	m.RLock()
	defer m.RUnlock()

	ctx, ok := m.context[r]
	return ctx, ok
}

func (m *ContextMap) Get(r *http.Request, k string) (interface{}, bool) {
	if ctx, ok := m.GetAll(r); ctx != nil {
		return ctx[k], ok
	}
	return nil, false
}

// Remove entries
func (m *ContextMap) RemoveAll(r *http.Request) {
	m.Lock()
	defer m.Unlock()

	delete(m.context, r)
}

func (m *ContextMap) Remove(r *http.Request, k string) {
	if ctx, _ := m.GetAll(r); ctx != nil {
		delete(ctx, k)
	}
}

// New Context Table
func NewContextMap() *ContextMap {
	return &ContextMap{
		context: make(map[*http.Request]Context),
	}
}
