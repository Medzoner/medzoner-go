package query

import (
	"context"
	"fmt"

	"github.com/Medzoner/medzoner-go/internal/domain/repository"

	"go.opentelemetry.io/otel/attribute"
	otelTrace "go.opentelemetry.io/otel/trace"
	"github.com/Medzoner/gomedz/pkg/observability"
)

// ListTechnoQueryHandler ListTechnoQueryHandler
type ListTechnoQueryHandler struct {
	TechnoRepository repository.TechnoRepository
}

// NewListTechnoQueryHandler NewListTechnoQueryHandler
func NewListTechnoQueryHandler(technoRepository repository.TechnoRepository) ListTechnoQueryHandler {
	return ListTechnoQueryHandler{
		TechnoRepository: technoRepository,
	}
}

// Handle handles ListTechnoQuery and return map[string]interface{}
func (l *ListTechnoQueryHandler) Handle(ctx context.Context, query ListTechnoQuery) (map[string]interface{}, error) {
	ctx, iSpan := observability.StartSpan(ctx, "ListTechnoQueryHandler.Publish",
		otelTrace.WithAttributes([]attribute.KeyValue{attribute.String("ctx", "request.Host")}...),
	)
	defer iSpan.End()

	resp := map[string]interface{}{}
	if query.Type == "stack" {
		resp, err := l.TechnoRepository.FetchStack(ctx)
		if err != nil {
			//return nil, fmt.Errorf("error fetching stack: %w", l.Telemetry.ErrorSpan(iSpan, err))
			return nil, fmt.Errorf("error fetching stack: %w", err)
		}
		return resp, nil
	}

	return resp, nil
}
