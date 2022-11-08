package main

import (
	"github.com/cyj19/gowalk"
	"github.com/cyj19/gowalk/logx"
)

type greeter struct {
	name string
}

func (g *greeter) Run() error {
	logx.Instance().Infof("hello, %s", g.name)
	return nil
}

func (g *greeter) Name() string {
	return "greeter"
}

var _ gowalk.Component = (*greeter)(nil)

func main() {
	_ = gowalk.Run(&greeter{name: "gowalk"})

	g, err := gowalk.GetComponent("greeter")
	if err != nil {
		logx.Instance().Fatal(err)
	}
	logx.Instance().Infof("greeter: %#v", g)
}
