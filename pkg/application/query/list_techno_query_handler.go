package query

import (
	"github.com/Medzoner/medzoner-go/pkg/domain/repository"
)

type ListTechnoQueryHandler struct {
	ItechnoRepository repository.TechnoRepository
}

func (l *ListTechnoQueryHandler) Handle(query ListTechnoQuery) map[string]interface{} {
	if query.Type == "stack" {
		return l.ItechnoRepository.FetchStack()
	}
	if query.Type == "experience" {
		return l.ItechnoRepository.FetchExperience()
	}
	if query.Type == "formation" {
		return l.ItechnoRepository.FetchFormation()
	}
	if query.Type == "lang" {
		return l.ItechnoRepository.FetchLang()
	}
	if query.Type == "other" {
		return l.ItechnoRepository.FetchOther()
	}
	return map[string]interface{}{}
}
