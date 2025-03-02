package main

import (
	"os"

	"github.com/baking-bread/boutique/internal/core/domain"
	"github.com/baking-bread/boutique/pkg/pagination"
	"github.com/baking-bread/boutique/pkg/plugin/shared"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type CatalogDemo struct {
	logger hclog.Logger
}

func (c *CatalogDemo) ListServices(p pagination.Pagination) (pagination.PagedResult[shared.CatalogItem], error) {
	c.logger.Info("ListServices called")
	return pagination.PagedResult[shared.CatalogItem]{}, nil
}

func (c *CatalogDemo) ListComponents(p pagination.Pagination) (pagination.PagedResult[shared.CatalogItem], error) {
	return pagination.PagedResult[shared.CatalogItem]{}, nil
}

func (c *CatalogDemo) ListTemplates(p pagination.Pagination) (pagination.PagedResult[shared.CatalogItem], error) {
	return pagination.PagedResult[shared.CatalogItem]{}, nil
}

func (c *CatalogDemo) GetService(id string) (domain.Service, error) {
	return domain.Service{}, nil
}

func (c *CatalogDemo) GetComponent(id string) (domain.Component, error) {
	return domain.Component{}, nil
}

func (c *CatalogDemo) GetTemplate(id string) (domain.Template, error) {
	return domain.Template{}, nil
}

func main() {

	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "demo-plugin",
		Output: os.Stderr,
		Level:  hclog.Trace,
	})

	demo := &CatalogDemo{
		logger: logger,
	}

	var pluginMap = map[string]plugin.Plugin{
		"catalog": &shared.CatalogPlugin{Impl: demo},
	}

	logger.Info("message from plugin", "foo", "bar")

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.HandshakeConfig,
		Plugins:         pluginMap,
		Logger:          logger,
	})
}
