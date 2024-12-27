package ports

import "github.com/baking-bread/boutique/internal/core/domain"

type Hub interface {
	GetCatalog(input GetCatalogInput) (out domain.Catalog, err error)
	ListCatalogs(input ListCatalogsInput) (out []domain.Catalog, err error)
	RegisterCatalog(input RegisterCatalogInput) (out domain.Catalog, err error)
	UnregisterCatalog(input UnregsiterCatalogInput) (out domain.Catalog, err error)
	UpdateCatalog(input UpdateCatalogInput) (out domain.Catalog, err error)
}

type CatalogRepository interface {
	GetCatalog(input GetCatalogInput) (out domain.Catalog, err error)
	ListCatalogs(input ListCatalogsInput) (out []domain.Catalog, err error)
	AddCatalog(input RegisterCatalogInput) (out domain.Catalog, err error)
	UpdateCatalog(input UpdateCatalogInput) (out domain.Catalog, err error)
	RemoveCatalog(input UnregsiterCatalogInput) (out domain.Catalog, err error)
}

type GetCatalogInput struct {
	domain.Catalog
}

type ListCatalogsInput struct {
	domain.Catalog
}

type RegisterCatalogInput struct {
	domain.Catalog
}

type UnregsiterCatalogInput struct {
	domain.Catalog
}

type UpdateCatalogInput struct {
	domain.Catalog
}
