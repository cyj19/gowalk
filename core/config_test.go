package core

import (
	"github.com/cyj19/gowalk/component/mysql"
	"testing"
)

func TestLoadComponentConfig(t *testing.T) {
	configPath := "./config_test.yml"
	err := LoadConfig(configPath)
	if err != nil {
		t.Error(err)
	}

	m := &mysql.Instance{}

	err = GetConfig(m.Name(), m)
	if err != nil {
		t.Error(err)
	}
	t.Logf("mysql: %#v \n", m)
}
