package web

import (
	"net/http"
)

type GetHandler interface {
	Get(http.ResponseWriter, *http.Request)
}

type HeadHandler interface {
	Head(http.ResponseWriter, *http.Request)
}

type PostHandler interface {
	Post(http.ResponseWriter, *http.Request)
}

type PutHandler interface {
	Put(http.ResponseWriter, *http.Request)
}

type DeleteHandler interface {
	Delete(http.ResponseWriter, *http.Request)
}

// MethodHandler
type methodHandler struct {
	methods []string
	handler []http.HandlerFunc
}

// MethodHandler as http.Handler
func (m methodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r != nil {
		for i, k := range m.methods {
			if k == r.Method {
				m.handler[i].ServeHTTP(w, r)
				return
			}
		}
	}

	// 405
	panic(MethodNotAllowed{
		methods: m.methods[:],
	})
}

// MethodHandler constructor
func (m *methodHandler) addMethodHandlerFunc(method string, h http.HandlerFunc) {
	m.methods = append(m.methods, method)
	m.handler = append(m.handler, h)
}

func MethodHandler(h interface{}) http.Handler {
	var m methodHandler
	var get GetHandler

	// GET
	if o, ok := h.(GetHandler); ok {
		get = o
		m.addMethodHandlerFunc(http.MethodGet, o.Get)
	}
	// HEAD
	if o, ok := h.(HeadHandler); ok {
		m.addMethodHandlerFunc(http.MethodHead, o.Head)
	} else if get != nil {
		m.addMethodHandlerFunc(http.MethodHead, get.Get)
	}
	// POST
	if o, ok := h.(PostHandler); ok {
		m.addMethodHandlerFunc(http.MethodPost, o.Post)
	}
	// PUT
	if o, ok := h.(PutHandler); ok {
		m.addMethodHandlerFunc(http.MethodPut, o.Put)
	}
	// DELETE
	if o, ok := h.(DeleteHandler); ok {
		m.addMethodHandlerFunc(http.MethodDelete o.Delete)
	}

	return m
}
