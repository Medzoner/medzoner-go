package query

import (
	"context"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/domain/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"go.opentelemetry.io/otel/attribute"
	otelTrace "go.opentelemetry.io/otel/trace"
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
	ctx, iSpan := l.Tracer.Start(ctx, "ListTechnoQueryHandler.Publish",
		otelTrace.WithAttributes([]attribute.KeyValue{attribute.String("ctx", "request.Host")}...),
	)
	defer iSpan.End()

	resp := map[string]interface{}{}
	if query.Type == "stack" {
		resp, err := l.TechnoRepository.FetchStack(ctx)
		if err != nil {
			return nil, l.Tracer.Error(iSpan, fmt.Errorf("error fetching stack: %w", err))
		}
		return resp, nil
	}

	return resp, nil
}
