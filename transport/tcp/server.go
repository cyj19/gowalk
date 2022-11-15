package tcp

import (
	"context"
	"github.com/cyj19/gowalk/logk"
	"github.com/cyj19/gowalk/transport"
	"net"
)

type ServerOption func(*Server)

func RunFunc(f RunFunction) ServerOption {
	return func(s *Server) {
		s.runFunc = f
	}
}

type RunFunction func(context.Context, net.Listener) error

// Server tcp服务
type Server struct {
	lis     net.Listener
	runFunc RunFunction
}

func NewServer(addr string, opts ...ServerOption) (*Server, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	s := &Server{lis: lis}

	for _, opt := range opts {
		opt(s)
	}

	return s, nil

}

func (s *Server) Start(ctx context.Context) error {
	logk.Infof("[TCP Server listening on: %s]", s.lis.Addr())
	return s.runFunc(ctx, s.lis)
}

func (s *Server) Stop(ctx context.Context) error {
	if done := ctx.Done(); done != nil {
		<-done
	}
	logk.Infof("[TCP Server stopping: %s]", s.lis.Addr())
	return s.lis.Close()
}

var _ transport.Server = (*Server)(nil)
