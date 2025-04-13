package query

import (
	"context"
	"fmt"

	"github.com/Medzoner/medzoner-go/pkg/domain/repository"
	"github.com/Medzoner/medzoner-go/pkg/infra/telemetry"
	"go.opentelemetry.io/otel/attribute"
	otelTrace "go.opentelemetry.io/otel/trace"
)

// ListTechnoQueryHandler ListTechnoQueryHandler
type ListTechnoQueryHandler struct {
	TechnoRepository repository.TechnoRepository
	Telemetry        telemetry.Telemeter
}

// NewListTechnoQueryHandler NewListTechnoQueryHandler
func NewListTechnoQueryHandler(technoRepository repository.TechnoRepository, tm telemetry.Telemeter) ListTechnoQueryHandler {
	return ListTechnoQueryHandler{
		TechnoRepository: technoRepository,
		Telemetry:        tm,
	}
}

// Handle handles ListTechnoQuery and return map[string]interface{}
func (l *ListTechnoQueryHandler) Handle(ctx context.Context, query ListTechnoQuery) (map[string]interface{}, error) {
	ctx, iSpan := l.Telemetry.Start(ctx, "ListTechnoQueryHandler.Publish",
		otelTrace.WithAttributes([]attribute.KeyValue{attribute.String("ctx", "request.Host")}...),
	)
	defer iSpan.End()

	resp := map[string]interface{}{}
	if query.Type == "stack" {
		resp, err := l.TechnoRepository.FetchStack(ctx)
		if err != nil {
			return nil, fmt.Errorf("error fetching stack: %w", l.Telemetry.ErrorSpan(iSpan, err))
		}
		return resp, nil
	}

	return resp, nil
}
