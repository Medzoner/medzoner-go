package query

import (
	"context"
	"github.com/Medzoner/medzoner-go/pkg/domain/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
)

// ListTechnoQueryHandler ListTechnoQueryHandler
type ListTechnoQueryHandler struct {
	TechnoRepository repository.TechnoRepository
	Tracer           tracer.Tracer
}

// NewListTechnoQueryHandler NewListTechnoQueryHandler
func NewListTechnoQueryHandler(technoRepository repository.TechnoRepository, tracer tracer.Tracer) ListTechnoQueryHandler {
	return ListTechnoQueryHandler{
		TechnoRepository: technoRepository,
		Tracer:           tracer,
	}
}

// Handle handles ListTechnoQuery and return map[string]interface{}
func (l *ListTechnoQueryHandler) Handle(ctx context.Context, query ListTechnoQuery) (map[string]interface{}, error) {
	_, iSpan := l.Tracer.Start(ctx, "ListTechnoQueryHandler.Publish")
	defer iSpan.End()

	resp := map[string]interface{}{}
	if query.Type == "stack" {
		resp, err := l.TechnoRepository.FetchStack()
		if err != nil {
			iSpan.RecordError(err)
			return nil, err
		}
		return resp, nil
	}

	return resp, nil
}
