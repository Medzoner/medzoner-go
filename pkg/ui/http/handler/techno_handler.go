package handler

import (
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"net/http"
)

//TechnoHandler TechnoHandler
type TechnoHandler struct {
	Template               templater.Templater
	ListTechnoQueryHandler query.ListTechnoQueryHandler
}

//TechnoView TechnoView
type TechnoView struct {
	Locale      string
	PageTitle   string
	Stacks      interface{}
	Experiences interface{}
	Formations  interface{}
	Langs       interface{}
	Others      interface{}
	TorHost     string
}

//IndexHandle IndexHandle
func (h *TechnoHandler) IndexHandle(w http.ResponseWriter, r *http.Request) {
	view := TechnoView{
		Locale:      "fr",
		PageTitle:   "MedZoner.com",
		Stacks:      h.ListTechnoQueryHandler.Handle(query.ListTechnoQuery{Type: "stack"}),
		Experiences: h.ListTechnoQueryHandler.Handle(query.ListTechnoQuery{Type: "experience"}),
		Formations:  h.ListTechnoQueryHandler.Handle(query.ListTechnoQuery{Type: "formation"}),
		Langs:       h.ListTechnoQueryHandler.Handle(query.ListTechnoQuery{Type: "lang"}),
		Others:      h.ListTechnoQueryHandler.Handle(query.ListTechnoQuery{Type: "other"}),
	}
	view.TorHost = r.Header.Get("TOR-HOST")
	_, err := h.Template.Render("technos", view, w, http.StatusOK)
	if err != nil {
		panic(err)
	}
	_ = r
}
