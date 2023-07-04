package query

import (
	"github.com/Medzoner/medzoner-go/pkg/domain/repository"
)

// ListTechnoQueryHandler ListTechnoQueryHandler
type ListTechnoQueryHandler struct {
	TechnoRepository repository.TechnoRepository
}

// Handle handles ListTechnoQuery and return map[string]interface{}
func (l *ListTechnoQueryHandler) Handle(query ListTechnoQuery) map[string]interface{} {
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
	return map[string]interface{}{}
}
