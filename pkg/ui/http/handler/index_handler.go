package handler

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"net/http"
)

//IndexHandler IndexHandler
type IndexHandler struct {
	Template               templater.Templater
	ListTechnoQueryHandler query.ListTechnoQueryHandler
	RecaptchaSiteKey       string
}

//IndexView IndexView
type IndexView struct {
	Locale    string
	PageTitle string
	TorHost   string
	TechnoView
	Message          interface{}
	Errors           interface{}
	RecaptchaSiteKey string
	PageDescription  string
}

//IndexHandle IndexHandle
func (h *IndexHandler) IndexHandle(w http.ResponseWriter, r *http.Request) {
	view := IndexView{
		Locale:    "fr",
		PageTitle: "MedZoner.com",
		TorHost:   r.Header.Get("TOR-HOST"),
		TechnoView: TechnoView{
			Stacks: h.ListTechnoQueryHandler.Handle(query.ListTechnoQuery{Type: "stack"}),
		},
		RecaptchaSiteKey: "6LeisDMhAAAAAPVU3SExc1nzrIUlZvL9SIKfMMye",
		PageDescription:  "Mehdi YOUB - Développeur Web Full Stack - NestJS Symfony Golang VueJS",
	}
	_, err := h.Template.Render("index", view, w, http.StatusOK)
	if err != nil {
		fmt.Println(err)
	}
	_ = r
}
