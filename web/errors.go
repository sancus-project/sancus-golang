package web

type MethodNotAllowed interface {
	Methods() []string
	Error() string
}
