package query

import (
	"context"
	"fmt"
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
func (l *ListTechnoQueryHandler) Handle(ctx context.Context, query ListTechnoQuery) map[string]interface{} {
	_, iSpan := l.Tracer.Start(ctx, fmt.Sprintf("ListTechnoQueryHandler.Handle"))
	iSpan.AddEvent("ListTechnoQueryHandler.Handle-Event")
	if query.Type == "stack" {
		return l.TechnoRepository.FetchStack()
	}
	if query.Type == "experience" {
		return l.TechnoRepository.FetchExperience()
	}
	if query.Type == "formation" {
		return l.TechnoRepository.FetchFormation()
	}
	if query.Type == "lang" {
		return l.TechnoRepository.FetchLang()
	}
	if query.Type == "other" {
		return l.TechnoRepository.FetchOther()
	}
	iSpan.End()
	return map[string]interface{}{}
}
