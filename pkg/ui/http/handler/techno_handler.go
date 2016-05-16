package handler

import (
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"html/template"
	"log"
	"net/http"
	"os"
)

type TechnoHandler struct {
	ListTechnoQueryHandler query.ListTechnoQueryHandler
}

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
	t := template.New("techno template")
	pp, _ := os.Getwd()
	t = template.Must(t.ParseFiles(
		pp+"/tmpl/base.html",
		pp+"/tmpl/footer.html",
		pp+"/tmpl/header.html",
		pp+"/tmpl/Technos/index.html",
		pp+"/tmpl/Technos/stack.html",
		pp+"/tmpl/Technos/experience.html",
		pp+"/tmpl/Technos/formation.html",
		pp+"/tmpl/Technos/lang.html",
		pp+"/tmpl/Technos/other.html",
	))
	w.WriteHeader(http.StatusOK)
	err := t.ExecuteTemplate(w, "layout", view)

	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
	_ = r
}
