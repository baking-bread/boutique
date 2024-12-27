package engine

import (
	"bytes"
	"errors"
	"html/template"
)

// Engine is the Boutique rendering implementation for services
type Engine struct {
	// If strict is enabled, template rendering will fail if a template references a value was missing
	Strict bool
}

// New creates a new instance of Engine
func New() Engine {
	return Engine{}
}

func (engine Engine) Parse(tpls map[string]string) (parsed *template.Template, err error) {
	return engine.parse(tpls)
}

func (engine Engine) Render(tpls map[string]string, values map[string]interface{}) (rendered map[string]string, err error) {
	root, err := engine.parse(tpls)
	if err != nil {
		return
	}

	rendered, err = engine.render(root, values)

	return
}

func (engine Engine) parse(tpls map[string]string) (parsed *template.Template, err error) {
	var errs []error
	parsed = template.New("root")

	for name, tpl := range tpls {
		_, parseError := parsed.New(name).Parse(tpl)
		if parseError != nil {
			errs = append(errs, parseError)
		}
	}

	err = errors.Join(errs...)

	return
}

func (engine Engine) render(root *template.Template, values map[string]interface{}) (rendered map[string]string, err error) {
	if engine.Strict {
		root.Option("missingkey=error")
	} else {
		root.Option("missingkey=zero")
	}

	tpls := root.Templates()
	rendered = make(map[string]string, len(tpls)-1)

	for _, tpl := range tpls {
		name := tpl.Name()

		// The root template is not rendered
		// It only serves as a container to share a context
		if root.Name() == name {
			continue
		}

		var buf bytes.Buffer
		execErr := root.ExecuteTemplate(&buf, name, values)
		if execErr != nil {
			errors.Join(err, execErr)
		}

		// TODO: check if the "<no value>" can be safely replaced by empty string ""
		rendered[name] = buf.String()
	}

	return
}
