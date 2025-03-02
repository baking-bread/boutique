package engine

import (
	"testing"
)

type Function struct{}

func (f Function) World() string {
	return "world"
}

func TestTemplate(t *testing.T) {
	var tpls map[string]string = make(map[string]string)

	var vals = make(map[string]any)

	tpls["one"] = "{{ .test }}"
	tpls["two"] = `{{ .functions.World }}`

	vals["test"] = "hello"
	vals["functions"] = Function{}

	engine := New()

	rendered, err := engine.Render(tpls, vals)

	if err != nil {
		t.Log(err.Error())
	}

	for key, value := range rendered {
		t.Log(key, value)
	}

}
