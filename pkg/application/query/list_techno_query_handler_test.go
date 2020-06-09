package query_test

import (
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"testing"
)

func TestListTechnoQueryHandler(t *testing.T) {

	t.Run("Unit: test ListTechnoQueryHandler \"stack\" success", func(t *testing.T) {
		listTechnoQuery := query.ListTechnoQuery{
			Type: "stack",
		}

		handler := query.ListTechnoQueryHandler{
			TechnoRepository: &TechnoRepositoryTest{},
		}
		handler.Handle(listTechnoQuery)
	})

	t.Run("Unit: test ListTechnoQueryHandler \"experience\" success", func(t *testing.T) {
		listTechnoQuery := query.ListTechnoQuery{
			Type: "experience",
		}
		handler := query.ListTechnoQueryHandler{
			TechnoRepository: &TechnoRepositoryTest{},
		}
		handler.Handle(listTechnoQuery)
	})

	t.Run("Unit: test ListTechnoQueryHandler \"formation\" success", func(t *testing.T) {
		listTechnoQuery := query.ListTechnoQuery{
			Type: "formation",
		}
		handler := query.ListTechnoQueryHandler{
			TechnoRepository: &TechnoRepositoryTest{},
		}
		handler.Handle(listTechnoQuery)
	})

	t.Run("Unit: test ListTechnoQueryHandler \"lang\" success", func(t *testing.T) {
		listTechnoQuery := query.ListTechnoQuery{
			Type: "lang",
		}
		handler := query.ListTechnoQueryHandler{
			TechnoRepository: &TechnoRepositoryTest{},
		}
		handler.Handle(listTechnoQuery)
	})

	t.Run("Unit: test ListTechnoQueryHandler \"other\" success", func(t *testing.T) {
		listTechnoQuery := query.ListTechnoQuery{
			Type: "other",
		}

		handler := query.ListTechnoQueryHandler{
			TechnoRepository: &TechnoRepositoryTest{},
		}
		handler.Handle(listTechnoQuery)
	})

	t.Run("Unit: test ListTechnoQueryHandler non existent type success", func(t *testing.T) {
		listTechnoQuery := query.ListTechnoQuery{
			Type: "fake",
		}

		handler := query.ListTechnoQueryHandler{
			TechnoRepository: &TechnoRepositoryTest{},
		}
		handler.Handle(listTechnoQuery)
	})

}

type TechnoRepositoryTest struct{}

func (m *TechnoRepositoryTest) FetchStack() map[string]interface{} {
	return map[string]interface{}{}
}
func (m *TechnoRepositoryTest) FetchExperience() map[string]interface{} {
	return map[string]interface{}{}
}
func (m *TechnoRepositoryTest) FetchFormation() map[string]interface{} {
	return map[string]interface{}{}
}
func (m *TechnoRepositoryTest) FetchLang() map[string]interface{} {
	return map[string]interface{}{}
}
func (m *TechnoRepositoryTest) FetchOther() map[string]interface{} {
	return map[string]interface{}{}
}
