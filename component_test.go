package gowalk

import (
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
	Run(&greeter{})
}
