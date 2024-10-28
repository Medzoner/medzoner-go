package handler

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/http_utils"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"net/http"
)

// NotFoundView NotFoundView
type NotFoundView struct {
	Locale          string
	PageTitle       string
	TorHost         string
	PageDescription string
}

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

// Handle handles NotFoundHandler
func (h *NotFoundHandler) Handle(w http.ResponseWriter, r *http.Request) {
	_, span := h.Tracer.StartRoot(r.Context(), r, "NotFoundHandler.Handle")
	defer span.End()

	view := &NotFoundView{
		Locale:          "fr",
		PageTitle:       "MedZoner.com - Not Found",
		TorHost:         r.Header.Get("TOR-HOST"),
		PageDescription: "MedZoner.com - Not Found",
	}
	_, err := h.Template.Render("404", view, w)
	if err != nil {
		http_utils.ResponseError(w, err, http.StatusInternalServerError, span)
		return
	}
}
