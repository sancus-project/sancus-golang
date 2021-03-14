package web

import (
	"go.sancus.io/core/net"
	"net/http"
)

// `net/http`.Server wrapper
// to use if `sancus/net`.Listen() is wanted
type Server struct {
	http.Server
}

func NewServer(addr string, h http.Handler) *Server {
	return &Server{Server: http.Server{
		Addr:    addr,
		Handler: h,
	}}
}

func (s *Server) ListenAndServe() error {
	addr := s.Addr
	if addr == "" {
		addr = ":http"
	}
	l, e := net.Listen(addr)
	if e != nil {
		return e
	}
	return s.Serve(l)
}
