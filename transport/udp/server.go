package udp

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

type RunFunction func(context.Context, *net.UDPConn) error

type Server struct {
	ip      string
	port    int
	lis     *net.UDPConn
	runFunc RunFunction
}

func NewServer(ip string, port int, opts ...ServerOption) (*Server, error) {
	lis, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IP(ip), Port: port})
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
	logk.GetLogger().Infof("[UDP Server listening on: %s]", s.lis.LocalAddr())
	return s.runFunc(ctx, s.lis)
}

func (s *Server) Stop(ctx context.Context) error {
	if done := ctx.Done(); done != nil {
		<-done
	}
	logk.GetLogger().Infof("[UDP Server stopping: %s]", s.lis.LocalAddr())
	return s.lis.Close()
}

var _ transport.Server = (*Server)(nil)
