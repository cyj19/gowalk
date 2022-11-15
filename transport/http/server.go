package http

import (
	"context"
	"crypto/tls"
	"github.com/cyj19/gowalk/logk"
	"github.com/cyj19/gowalk/transport"
	"net"
	"net/http"
	"time"
)

type ServerOption func(*Server)

func Handler(handler http.Handler) ServerOption {
	return func(s *Server) {
		s.srv.Handler = handler
	}
}

func ReadTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.srv.ReadTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.srv.WriteTimeout = timeout
	}
}

func IdleTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.srv.IdleTimeout = timeout
	}
}

func TLSConfig(cfg *tls.Config) ServerOption {
	return func(s *Server) {
		s.srv.TLSConfig = cfg
	}
}

// Server http 服务
type Server struct {
	srv *http.Server
}

func NewServer(addr string, opts ...ServerOption) *Server {
	srv := &Server{srv: &http.Server{Addr: addr, Handler: http.DefaultServeMux}}
	for _, opt := range opts {
		opt(srv)
	}
	return srv
}

func (s *Server) Start(ctx context.Context) error {
	// 服务的根上下文
	s.srv.BaseContext = func(listener net.Listener) context.Context {
		return ctx
	}
	logk.Infof("[HTTP] server listening on: %s", s.srv.Addr)
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	logk.Infof("[HTTP] server stopping: %s", s.srv.Addr)
	return s.srv.Shutdown(ctx)
}

var _ transport.Server = (*Server)(nil)
