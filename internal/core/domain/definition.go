package domain

type TemplateDefinition struct {
	Source       TemplateSource
	Values       []TemplateValue
	Dependencies []TemplateDependency
	Policies     []TemplatePolicy
}
