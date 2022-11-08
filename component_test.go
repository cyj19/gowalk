package gowalk

import (
	"github.com/cyj19/gowalk/config"
	"log"
	"testing"
)

type greeter struct {
}

func (g *greeter) Run() error {
	log.Println("loading greeter...")
	return nil
}

func (g *greeter) Name() string {
	return "greeter"
}

func TestAddAndLoadComponents(t *testing.T) {
	_ = config.LoadConfig("./config.dev.yml")
	_ = AddAndLoadComponents(&greeter{})
}
