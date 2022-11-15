package udp

import (
	"context"
	"github.com/cyj19/gowalk/logk"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	logk.SetupLog("./", logk.LogConfig{})
	s, err := NewServer("127.0.0.1", 8888, RunFunc(udp))
	if err != nil {
		t.Fatal(err)
	}

	if err = s.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func udp(ctx context.Context, conn *net.UDPConn) error {
	// TODO...

	return nil
}
