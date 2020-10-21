package handler

import (
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"net/http"
)

type IndexHandler struct {
	Template templater.Templater
}

type IndexView struct {
	Locale    string
	PageTitle string
	TorHost   string
}

func (h *IndexHandler) IndexHandle(w http.ResponseWriter, r *http.Request) {
	view := IndexView{
		Locale: "fr",
		PageTitle: "MedZoner.com",
	}
	view.TorHost = r.Header.Get("TOR-HOST")
	_, err := h.Template.Render("index", view, w, http.StatusOK)
	if err != nil {
		panic(err)
	}
	_ = r
}
