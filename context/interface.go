package context

import (
	"net/http"
)

type Context map[string]interface{}

type RequestContextMapper interface {
	GetAll(r *http.Request) Context
	Get(r *http.Request, k string) (interface{}, bool)
	Set(r *http.Request, k string, v interface{})
	RemoveAll(r *http.Request)
	Remove(r *http.Request, k string)
}
