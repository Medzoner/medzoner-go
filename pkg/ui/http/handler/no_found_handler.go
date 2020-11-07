package handler

import (
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"net/http"
)

//NotFoundHandler NotFoundHandler
type NotFoundHandler struct {
	Template templater.Templater
}

//Handle Handle
func (h *NotFoundHandler) Handle(w http.ResponseWriter, r *http.Request) {
	_, err := h.Template.Render("404", nil, w, http.StatusNotFound)
	if err != nil {
		panic(err)
	}
	_ = r
}
