package tcp

import (
	"context"
	"github.com/cyj19/gowalk/logk"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	logk.SetupLog("./", logk.LogConfig{})

	s, err := NewServer(":8888", RunFunc(tcp))
	if err != nil {
		t.Fatal(err)
	}
	if err = s.Start(context.Background()); err != nil {
		t.Fatal(err)
	}

}

func tcp(ctx context.Context, lis net.Listener) error {
	// TODO...
	for {
		conn, err := lis.Accept()
		if err != nil {
			return err
		}
		ch := make(chan []byte, 1024)
		go send(conn, ch)
		go read(conn, ch)
	}
	return nil
}

func send(conn net.Conn, ch <-chan []byte) {
	msg := <-ch
	conn.Write(msg)
}

func read(conn net.Conn, ch chan<- []byte) {
	msg := make([]byte, 0, 1024)
	conn.Read(msg)
	ch <- msg
}
