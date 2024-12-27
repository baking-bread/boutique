package cataloghub

import (
	"github.com/baking-bread/boutique/internal/core/domain"
	"github.com/baking-bread/boutique/internal/core/ports"
)

type Hub struct {
	CatalogRepository ports.CatalogRepository
}

func New(catalogRepository ports.CatalogRepository) *Hub {
	return &Hub{
		CatalogRepository: catalogRepository,
	}
}

func (hub Hub) GetCatalog(input ports.GetCatalogInput) (out domain.Catalog, err error) {
	out, err = hub.CatalogRepository.GetCatalog(input)
	return out, err
}

func (hub Hub) ListCatalogs(input ports.ListCatalogsInput) (out []domain.Catalog, err error) {
	out, err = hub.CatalogRepository.ListCatalogs(input)
	return out, err
}

func (hub Hub) RegisterCatalog(input ports.RegisterCatalogInput) (out domain.Catalog, err error) {
	out, err = hub.CatalogRepository.AddCatalog(input)
	return out, err
}

func (hub Hub) UnregisterCatalog(input ports.UnregsiterCatalogInput) (out domain.Catalog, err error) {
	out, err = hub.CatalogRepository.RemoveCatalog(input)
	return out, err
}

func (hub Hub) UpdateCatalog(input ports.UpdateCatalogInput) (out domain.Catalog, err error) {
	out, err = hub.CatalogRepository.UpdateCatalog(input)
	return out, err
}
