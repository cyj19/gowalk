package core

import (
	"github.com/cyj19/gowalk/component/helloworld"
	"testing"
)

func TestLoadComponents(t *testing.T) {
	_ = AddAndLoadComponents(&helloworld.Instance{Message: "hawk"})
	t.Logf("components: %#v \n", components)

	h, err := GetComponent("helloworld")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v \n", h)
}
