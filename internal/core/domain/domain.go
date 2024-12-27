package domain

// Contains a collection of services, bundles and or packages
type Catalog struct {
	Id       string
	Brokers  []Broker
	Services []Service
}

// Services offer actions
type Service struct {
	Id          string
	Catalog     Catalog
	Path        string // item location within catalog
	Ref         string // for versioning
	Description string
	Output      map[string]ServiceOutput // outputs from this service
}

type ServiceOutput struct {
	Name        string
	Description string
	Type        string
}

// Bundles customizes a set of services
// By structuring a service hierarchy adding policies and more
type Bundle struct {
	Id        string
	Path      string
	Ref       string
	Config    BundleConfig
	Templates []TemplateContent
	// policies
}

type BundleConfig struct {
	Description string
	Tags        []string
}

type TemplateContent struct {
	Name    string
	Content string
}

// Brokers can validate bundle templates and execute services
type Broker struct {
	Id          string
	Name        string
	Description string
}

// Packages describe results after a bundle was personalized
type Package struct {
	Id         string
	Catalog    Catalog
	Path       string
	Ref        string
	TemplateId string
	Instances  []Instance

	// history
	// status
}

type PackageConfig struct {
	Description string
	Tags        []string
}

type Instance struct {
	Name    string
	Content string
}
