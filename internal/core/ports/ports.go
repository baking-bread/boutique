package ports

type PolicyService interface{}
type ValueService interface{}
type DependencyService interface{}

type LifecycleService interface {
	Preview()
	Create()
	Get()
	List()
	Update()
	Delete()
	Status()
}

type DefinitionService interface {
	List()
	Get()
	Add()
	Remove()
}

type TemplateRepository interface {
	List()
	Save()
	Read()
}
