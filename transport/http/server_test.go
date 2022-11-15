package http

import (
	"context"
	"github.com/cyj19/gowalk/logk"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	logk.SetupLog("./", logk.LogConfig{})
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		logk.Info(r.URL.Path)
		w.Write([]byte("pong"))
	})
	s := NewServer(":8888", Handler(mux))

	if err := s.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
}
