package handler

import (
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"net/http"
)

//NotFoundHandler NotFoundHandler
type NotFoundHandler struct {
	Template templater.Templater
}

//NotFoundView NotFoundView
type NotFoundView struct {
	Locale    string
	PageTitle string
	TorHost   string
}

//Handle Handle
func (h *NotFoundHandler) Handle(w http.ResponseWriter, r *http.Request) {
	view := &NotFoundView{
		Locale:    "fr",
		PageTitle: "MedZoner.com - Not Found",
		TorHost: r.Header.Get("TOR-HOST"),
	}
	_, err := h.Template.Render("404", view, w, http.StatusNotFound)
	if err != nil {
		panic(err)
	}
	_ = r
}
