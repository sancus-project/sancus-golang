package context

import (
	"net/http"
	"sync"
)

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
func (m *ContextMap) GetAll(r *http.Request) Context {
	m.RLock()
	defer m.RUnlock()

	return m.context[r]
}

func (m *ContextMap) Get(r *http.Request, k string) (interface{}, bool) {
	if ctx := m.GetAll(r); ctx != nil {
		v, ok := ctx[k]
		return v, ok
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
	if ctx := m.GetAll(r); ctx != nil {
		delete(ctx, k)
	}
}

// New Context Table
func NewContextMap() *ContextMap {
	return &ContextMap{
		context: make(map[*http.Request]Context),
	}
}
