package gowalk

import (
	"context"
	"github.com/cyj19/gowalk/transport"
	"os"
	"time"
)

type Option func(o *option)

type option struct {
	ctx         context.Context
	id          string
	name        string
	version     string
	metadata    map[string]interface{}
	sigs        []os.Signal
	servers     []transport.Server
	stopTimeout time.Duration
}

func Context(ctx context.Context) Option {
	return func(o *option) {
		o.ctx = ctx
	}
}

func ID(id string) Option {
	return func(o *option) {
		o.id = id
	}
}

func Name(name string) Option {
	return func(o *option) {
		o.name = name
	}
}

func Version(version string) Option {
	return func(o *option) {
		o.version = version
	}
}

func Metadata(md map[string]interface{}) Option {
	return func(o *option) {
		o.metadata = md
	}
}

func Signal(sigs ...os.Signal) Option {
	return func(o *option) {
		o.sigs = sigs
	}
}

func Servers(srv ...transport.Server) Option {
	return func(o *option) {
		o.servers = srv
	}
}

func StopTimeout(t time.Duration) Option {
	return func(o *option) {
		o.stopTimeout = t
	}
}
