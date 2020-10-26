package handler

import (
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"net/http"
)

//IndexHandler IndexHandler
type IndexHandler struct {
	Template templater.Templater
}

//IndexView IndexView
type IndexView struct {
	Locale    string
	PageTitle string
	TorHost   string
}

//IndexHandle IndexHandle
func (h *IndexHandler) IndexHandle(w http.ResponseWriter, r *http.Request) {
	view := IndexView{
		Locale:    "fr",
		PageTitle: "MedZoner.com",
	}
	view.TorHost = r.Header.Get("TOR-HOST")
	_, err := h.Template.Render("index", view, w, http.StatusOK)
	if err != nil {
		panic(err)
	}
	_ = r
}
