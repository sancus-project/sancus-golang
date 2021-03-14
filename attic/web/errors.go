package web

import (
	"fmt"
	"net/http"
	"strings"
)

type HttpError interface {
	error
	fmt.Stringer
	http.Handler

	Code() int
}

// 405
type MethodNotAllowed struct {
	HttpError
	methods []string
}

func (m MethodNotAllowed) Methods() []string {
	return m.methods
}

// MethodNotAllowed as http.Handler
func (m MethodNotAllowed) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", strings.Join(m.methods, ", "))
	http.Error(w, fmt.Sprintf("Method %s not allowed", r.Method), m.Code())
}

// MethodNotAllowed as HttpError
func (m MethodNotAllowed) Code() int {
	return http.StatusMethodNotAllowed
}

// MethodNotAllowed as fmt.Stringer
func (m MethodNotAllowed) String() string {
	return "Method not allowed"
}

// MethodNotAllowed as error
func (m MethodNotAllowed) Error() string {
	return fmt.Sprintf("Allow: %s", strings.Join(m.methods, ", "))
}
