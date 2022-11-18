package gowalk

import (
	"github.com/cyj19/gowalk/logk"
	"testing"
)

type greeter struct {
}

func (g *greeter) Run() error {
	logk.GetLogger().Info("loading greeter...")
	return nil
}

func (g *greeter) Name() string {
	return "greeter"
}

func TestAddAndLoadComponents(t *testing.T) {
	_ = AddAndLoadComponents(&greeter{})
}
