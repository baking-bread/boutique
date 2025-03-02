package services

import (
	"github.com/baking-bread/boutique/internal/core/domain"
	"github.com/baking-bread/boutique/pkg/pagination"
)

type TemplateExecutionService struct{}

func NewTemplateExecutionService() *TemplateExecutionService {
	return &TemplateExecutionService{}
}

func (s *TemplateExecutionService) GetTemplateRun(id string) (run *domain.TemplateRun, err error) {
	return nil, nil
}

func (s *TemplateExecutionService) ListTemplateRuns(p pagination.Pagination) (pagination.PagedResult[domain.TemplateRun], error) {
	return pagination.PagedResult[domain.TemplateRun]{}, nil
}

func (s *TemplateExecutionService) PrepareTemplate(t *domain.Template) (run *domain.TemplateRun, err error) {
	return nil, nil
}

func (s *TemplateExecutionService) ExecuteTemplateRun(run *domain.TemplateRun) (err error) {
	return nil
}

func (s *TemplateExecutionService) CancelTemplateRun(id string) (t *domain.Template, err error) {
	return nil, nil
}

func (s *TemplateExecutionService) CreateTemplateRun(t *domain.Template, inputs []domain.Input) (run *domain.TemplateRun, err error) {
	return nil, nil
}

func (s *TemplateExecutionService) CreateServiceRun(t *domain.TemplateRun, step domain.Step, inputs []domain.Input) (run *domain.ServiceRun, err error) {
	return nil, nil
}
