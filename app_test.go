package gowalk

import (
	"github.com/cyj19/gowalk/logk"
	khttp "github.com/cyj19/gowalk/transport/http"
	"net/http"
	"testing"
)

func TestApp(t *testing.T) {

	// http服务：8888
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		logk.Info(r.URL.Path)
		w.Write([]byte("pong"))
	})
	hs := khttp.NewServer(":8888", khttp.Handler(mux))

	// http服务：8889
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		logk.Info(r.URL.Path)
		w.Write([]byte("hello world"))
	})

	hs2 := khttp.NewServer(":8889", khttp.Handler(mux2))

	app := New(Servers(hs, hs2))

	if err := app.Run(); err != nil {
		logk.Fatal(err)
	}
}
