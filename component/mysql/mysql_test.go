package mysql

import (
	"github.com/cyj19/gowalk/core"
	"github.com/cyj19/gowalk/core/conf"
	"testing"
)

func TestLoadMySQL(t *testing.T) {
	err := conf.LoadConfig("../../core/config.dev.yml")
	if err != nil {
		t.Error(err)
	}
	_ = core.AddAndLoadComponents(&Instance{})

}
