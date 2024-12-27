package engine

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	var tpls map[string]string = make(map[string]string)

	tpls["one"] = "hello "
	tpls["two"] = "world"

	engine := New()

	rendered, err := engine.Render(tpls, nil)

	if err != nil {
		t.Log(err.Error())
	}

	for key, value := range rendered {
		t.Log(key, value)
	}

}
