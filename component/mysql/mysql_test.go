package mysql

import (
	"github.com/cyj19/gowalk/core"
	"testing"
)

func TestLoadMySQL(t *testing.T) {
	err := core.LoadConfig("../../core/config.dev.yml")
	if err != nil {
		t.Error(err)
	}
	_ = core.AddAndLoadComponents(&Instance{})

}
