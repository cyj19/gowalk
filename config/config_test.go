package config

import (
	"testing"
)

type hello struct {
	Items map[string]item `json:"items"`
}

type item struct {
	Name string
	Age  int
}

func TestLoadComponentConfig(t *testing.T) {
	configPath := "../config.yml"
	err := LoadConfig(configPath)
	if err != nil {
		t.Error(err)
	}

	h := &hello{}
	err = GetConfig("hello", h)
	if err != nil {
		t.Error(err)
	}
	t.Logf("hello: %#v \n", h)
}
