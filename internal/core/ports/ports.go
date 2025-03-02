package ports

import "github.com/baking-bread/boutique/internal/core/domain"

// Pagination defines the structure for paginated results
type Pagination struct {
	Offset int
	Limit  int
}

// PagedResult wraps paginated data with metadata
type PagedResult[T any] struct {
	Items      []T
	TotalCount int
	HasMore    bool
}

// driver ports

// AgentManager handles agent-related operations
type AgentManager interface {
	CreateAgent(agent domain.Agent) (domain.Agent, error)
	UpdateAgent(id string, agent domain.Agent) (domain.Agent, error)
	GetAgent(id string) (domain.Agent, error)
	DeleteAgent(id string) error
	ListAgents(pagination Pagination) (PagedResult[domain.Agent], error)
	SearchAgents(query string, pagination Pagination) (PagedResult[domain.Agent], error)
}

// ServiceManager handles service-related operations
type ServiceManager interface {
	CreateService(service domain.Service) (domain.Service, error)
	UpdateService(id string, service domain.Service) (domain.Service, error)
	GetService(id string) (domain.Service, error)
	DeleteService(id string) error
	ListServices(pagination Pagination) (PagedResult[domain.Service], error)
	SearchServices(query string, pagination Pagination) (PagedResult[domain.Service], error)
}

// ComponentManager handles component-related operations
type ComponentManager interface {
	CreateComponent(component domain.Component) (domain.Component, error)
	UpdateComponent(id string, component domain.Component) (domain.Component, error)
	GetComponent(id string) (domain.Component, error)
	DeleteComponent(id string) error
	ListComponents(pagination Pagination) (PagedResult[domain.Component], error)
	SearchComponents(query string, pagination Pagination) (PagedResult[domain.Component], error)
}

// PolicyManager handles policy-related operations
type PolicyManager interface {
	CreatePolicy(policy domain.Policy) (domain.Policy, error)
	UpdatePolicy(id string, policy domain.Policy) (domain.Policy, error)
	GetPolicy(id string) (domain.Policy, error)
	DeletePolicy(id string) error
	ListPolicies(pagination Pagination) (PagedResult[domain.Policy], error)
	SearchPolicies(query string, pagination Pagination) (PagedResult[domain.Policy], error)
}

// TemplateManager handles template-related operations
type TemplateManager interface {
	CreateTemplate(template domain.Template) (domain.Template, error)
	UpdateTemplate(id string, template domain.Template) (domain.Template, error)
	GetTemplate(id string) (domain.Template, error)
	DeleteTemplate(id string) error
	ListTemplates(pagination Pagination) (PagedResult[domain.Template], error)
	SearchTemplates(query string, pagination Pagination) (PagedResult[domain.Template], error)
}

// ServiceRunManager handles service execution operations
type ServiceRunManager interface {
	StartServiceRun(serviceId string, inputs []domain.Input) (domain.ServiceRun, error)
	GetServiceRun(id string) (domain.ServiceRun, error)
	ListServiceRuns(pagination Pagination) (PagedResult[domain.ServiceRun], error)
	SearchServiceRuns(query string, pagination Pagination) (PagedResult[domain.ServiceRun], error)
	CancelServiceRun(id string) error
	GetServiceRunStatus(id string) (string, error)
}

// TemplateRunManager handles template execution operations
type TemplateRunManager interface {
	StartTemplateRun(templateId string, inputs map[string][]domain.Input) (domain.TemplateRun, error)
	GetTemplateRun(id string) (domain.TemplateRun, error)
	ListTemplateRuns(pagination Pagination) (PagedResult[domain.TemplateRun], error)
	SearchTemplateRuns(query string, pagination Pagination) (PagedResult[domain.TemplateRun], error)
	CancelTemplateRun(id string) error
	GetTemplateRunStatus(id string) (string, error)
}

// InstanceManager handles instance-related operations
type InstanceManager interface {
	CreateInstance(instance domain.Instance) (domain.Instance, error)
	UpdateInstance(id string, instance domain.Instance) (domain.Instance, error)
	GetInstance(id string) (domain.Instance, error)
	DeleteInstance(id string) error
	ListInstances(pagination Pagination) (PagedResult[domain.Instance], error)
	SearchInstances(query string, pagination Pagination) (PagedResult[domain.Instance], error)
	UpdateAttributes(id string, attributes []any) error
}

// LinkManager handles relationships between instances and service runs
type LinkManager interface {
	CreateLink(link domain.Link) (domain.Link, error)
	GetLink(id string) (domain.Link, error)
	DeleteLink(id string) error
	ListLinks(instanceId string, pagination Pagination) (PagedResult[domain.Link], error)
	SearchLinks(query string, pagination Pagination) (PagedResult[domain.Link], error)
	ValidateLink(instanceId string, serviceRunId string) error
}

// OwnershipManager handles ownership-related operations
type OwnershipManager interface {
	TransferOwnership(resourceId string, newOwner string) error
	ValidateOwnership(resourceId string, ownerId string) error
	ListResourcesByOwner(ownerId string, resourceType string, pagination Pagination) (PagedResult[any], error)
	GetResourceOwner(resourceId string) (string, error)
}

// driven ports

// Agent Instructions
// Component Status
// Execute Steps
// Fetch data
