package shared

import (
	"net/rpc"

	"github.com/baking-bread/boutique/internal/core/domain"
	"github.com/baking-bread/boutique/pkg/pagination"
	"github.com/hashicorp/go-plugin"
)

type CatalogItem struct {
	ID       string
	Checksum string
}

type Catalog interface {
	ListServices(p pagination.Pagination) (pagination.PagedResult[CatalogItem], error)
	ListComponents(p pagination.Pagination) (pagination.PagedResult[CatalogItem], error)
	ListTemplates(p pagination.Pagination) (pagination.PagedResult[CatalogItem], error)

	GetService(id string) (domain.Service, error)
	GetComponent(id string) (domain.Component, error)
	GetTemplate(id string) (domain.Template, error)
}

// Client implementation

type CatalogRPCClient struct {
	client *rpc.Client
}

func (c *CatalogRPCClient) ListServices(p pagination.Pagination) (pagination.PagedResult[CatalogItem], error) {
	var result pagination.PagedResult[CatalogItem]
	err := c.client.Call("Plugin.ListServices", p, &result)
	return result, err
}

func (c *CatalogRPCClient) ListComponents(p pagination.Pagination) (pagination.PagedResult[CatalogItem], error) {
	var result pagination.PagedResult[CatalogItem]
	err := c.client.Call("Plugin.ListComponents", p, &result)
	return result, err
}

func (c *CatalogRPCClient) ListTemplates(p pagination.Pagination) (pagination.PagedResult[CatalogItem], error) {
	var result pagination.PagedResult[CatalogItem]
	err := c.client.Call("Plugin.ListTemplates", p, &result)
	return result, err
}

func (c *CatalogRPCClient) GetService(id string) (domain.Service, error) {
	var result domain.Service
	err := c.client.Call("Plugin.GetService", id, &result)
	return result, err
}

func (c *CatalogRPCClient) GetComponent(id string) (domain.Component, error) {
	var result domain.Component
	err := c.client.Call("Plugin.GetComponent", id, &result)
	return result, err
}

func (c *CatalogRPCClient) GetTemplate(id string) (domain.Template, error) {
	var result domain.Template
	err := c.client.Call("Plugin.GetTemplate", id, &result)
	return result, err
}

// Server implementation

type CatalogRPCServer struct {
	Impl Catalog
}

func (s *CatalogRPCServer) ListServices(p pagination.Pagination, result *pagination.PagedResult[CatalogItem]) (err error) {
	*result, err = s.Impl.ListServices(p)
	return err
}

func (s *CatalogRPCServer) ListComponents(p pagination.Pagination, result *pagination.PagedResult[CatalogItem]) (err error) {
	*result, err = s.Impl.ListComponents(p)
	return err
}

func (s *CatalogRPCServer) ListTemplates(p pagination.Pagination, result *pagination.PagedResult[CatalogItem]) (err error) {
	*result, err = s.Impl.ListTemplates(p)
	return err
}

func (s *CatalogRPCServer) GetService(id string, result *domain.Service) (err error) {
	*result, err = s.Impl.GetService(id)
	return err
}

func (s *CatalogRPCServer) GetComponent(id string, result *domain.Component) (err error) {
	*result, err = s.Impl.GetComponent(id)
	return err
}

func (s *CatalogRPCServer) GetTemplate(id string, result *domain.Template) (err error) {
	*result, err = s.Impl.GetTemplate(id)
	return err
}

// Plugin implementation
type CatalogPlugin struct {
	Impl Catalog
}

// Server should return the RPC server compatible struct to serve
// the methods that the Client calls over net/rpc.
func (p *CatalogPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &CatalogRPCServer{Impl: p.Impl}, nil
}

// Client returns an interface implementation for the plugin you're
// serving that communicates to the server end of the plugin.
func (CatalogPlugin) Client(m *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &CatalogRPCClient{client: c}, nil
}
