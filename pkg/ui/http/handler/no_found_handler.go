package handler

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"net/http"
)

// NotFoundHandler NotFoundHandler
type NotFoundHandler struct {
	Template templater.Templater
	Tracer   tracer.Tracer
}

// NewNotFoundHandler NewNotFoundHandler
func NewNotFoundHandler(template templater.Templater, tracer tracer.Tracer) *NotFoundHandler {
	return &NotFoundHandler{
		Template: template,
		Tracer:   tracer,
	}
}

// NotFoundView NotFoundView
type NotFoundView struct {
	Locale          string
	PageTitle       string
	TorHost         string
	PageDescription string
}

// Handle Handle
func (h *NotFoundHandler) Handle(w http.ResponseWriter, r *http.Request) {
	h.Tracer.WriteLog(r.Context(), "NotFoundHandle")
	view := &NotFoundView{
		Locale:          "fr",
		PageTitle:       "MedZoner.com - Not Found",
		TorHost:         r.Header.Get("TOR-HOST"),
		PageDescription: "MedZoner.com - Not Found",
	}
	_, err := h.Template.Render("404", view, w, http.StatusNotFound)
	if err != nil {
		panic(err)
	}
	_ = r
	return
}
