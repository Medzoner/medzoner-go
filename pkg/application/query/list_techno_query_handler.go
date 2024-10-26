package query

import (
	"context"
	"github.com/Medzoner/medzoner-go/pkg/domain/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/middleware"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"go.opentelemetry.io/otel/attribute"
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
func (l *ListTechnoQueryHandler) Handle(ctx context.Context, query ListTechnoQuery) map[string]interface{} {
	_, iSpan := l.Tracer.Start(ctx, "ListTechnoQueryHandler.Handle")
	defer func() {
		iSpan.End()
	}()
	correlationID := middleware.GetCorrelationID(ctx)
	iSpan.SetAttributes(attribute.String("correlation_id", correlationID))

	resp := map[string]interface{}{}
	if query.Type == "stack" {
		resp = l.TechnoRepository.FetchStack()
	}

	return resp
}
