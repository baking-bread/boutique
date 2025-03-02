package domain

// Service catalog definition

type Agent struct {
	Id          string
	Name        string
	Description string
	Owner       string
	Tags        []string
	Policies    []Policy
}

type Service struct {
	Id          string
	Checksum    string
	Agent       Agent
	Name        string
	Description string
	Owner       string
	Ref         string // for versioning
	Action      string // create, update, delete, migration, rollback
	Inputs      []InputSchema
	Outputs     []OutputSchema
	Policies    []Policy
	// Rollback Service?
}

type Policy struct {
	Id          string
	Checksum    string
	Name        string
	Type        string // apply to read, write, delete, execute
	Description string
}

type InputSchema struct {
	Id          string
	Name        string
	Description string
	Type        string     // string, number, boolean, object, array, component, data (means: fetching data from an agent)
	Default     any        // can be nil
	Component   Component  // has to be set if type is component
	Data        DataSchema // has to be set if type is data
	Policies    []Policy
}

type OutputSchema struct {
	Id          string
	Name        string
	Description string
	Type        string
	Component   Component
	Policies    []Policy
}

type AttributeSchema struct {
	Id          string
	Name        string
	Description string
	Type        string
	Optional    bool
	Policies    []Policy
}

type Component struct {
	Id          string
	Name        string
	Checksum    string
	Description string
	Ref         string // for versioning
	Agent       Agent
	Attributes  []AttributeSchema
	Policies    []Policy
	// Include a hierarchy?
}

type DataSchema struct {
	Id          string
	Name        string
	Description string
	Type        string
	Policies    []Policy
}

type InputMapping struct {
	Id          string
	InputSchema InputSchema
	Mapping     string // string template
}

type Step struct {
	Id            string
	Name          string
	Description   string
	Service       Service
	InputMappings []InputMapping
}

type Template struct {
	Id          string
	Name        string
	Checksum    string
	Description string
	Ref         string // for versioning
	Steps       []Step
	Policies    []Policy
}

// Plugins can extend catalog items with additional attributes
// The agent manages the value of the attributes
type Extension struct {
	Id          string
	Name        string
	Category    string // template, service, component, agent
	Description string
	Owner       string
	Ref         string // for versioning
	Attributes  []AttributeSchema
	Policies    []Policy
}

// Instances

type Input struct {
	Schema InputSchema
	Value  any
}

type Output struct {
	Schema OutputSchema
	Value  any
}

type Data struct {
	Schema DataSchema
	Value  any
}

type ServiceRun struct {
	Id       string
	Start    string
	End      string
	Agent    Agent
	Service  Service
	Issuer   string
	Owner    string
	Inputs   []Input
	Outputs  []Output
	Status   string
	Policies []Policy
}

type TemplateRun struct {
	Id       string
	Issuer   string
	Owner    string
	Steps    []ServiceRun
	Policies []Policy
}

type Link struct {
	Id       string
	Issuer   string
	Owner    string
	Instance Instance
	Service  ServiceRun
	Policies []Policy
}

type Instance struct {
	Id         string
	Component  Component
	Agent      Agent
	Issuer     string
	Owner      string
	Attributes []any
	Links      []Link
	Policies   []Policy
}
