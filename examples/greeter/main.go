package main

import (
	"github.com/cyj19/gowalk"
	"github.com/cyj19/gowalk/logk"
)

type greeter struct {
	name string
}

func (g *greeter) Run() error {
	logk.Infof("hello, %s", g.name)
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
		logk.Fatal(err)
	}
	logk.Infof("greeter: %#v", g)
}
