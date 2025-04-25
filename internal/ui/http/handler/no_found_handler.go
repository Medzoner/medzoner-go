package handler

import (
	"net/http"

	"github.com/Medzoner/medzoner-go/internal/ui/http/http_utils"
	"github.com/Medzoner/medzoner-go/internal/ui/http/templater"
	"github.com/Medzoner/medzoner-go/pkg/infra/telemetry"
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
	Tracer   telemetry.Telemeter
}

// NewNotFoundHandler NewNotFoundHandler
func NewNotFoundHandler(template templater.Templater, tracer telemetry.Telemeter) *NotFoundHandler {
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

	w.WriteHeader(http.StatusNotFound)

	if err := h.Template.Render("404", view, w); err != nil {
		http_utils.ResponseError(w, err, http.StatusInternalServerError, span)
		return
	}
}
