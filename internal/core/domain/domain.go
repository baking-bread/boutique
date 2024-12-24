package domain

import catalogtype "github.com/baking-bread/boutique/internal/core/domain/enum/catalog_type"

// Contains a collection of templates and or packages
type Catalog struct {
	Type catalogtype.CatalogType
}

type Chart struct {
	Catalog   Catalog
	Path      string // item location within catalog
	Ref       string // for versioning
	Config    TemplateConfiguration
	Templates []TemplateContent

	// policies
}

type TemplateConfiguration struct{}
type TemplateContent struct{}

type Package struct {
	Catalog   Catalog
	Path      string
	Ref       string
	Instances []InstanceContent

	// history
	// status
}

type InstanceContent struct{}
