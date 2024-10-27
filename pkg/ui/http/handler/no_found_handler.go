package handler

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	otelTrace "go.opentelemetry.io/otel/trace"
	"net/http"
	"time"

	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
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
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
	defer cancel()
	ctx, span := h.Tracer.Start(
		ctx,
		"IndexHandler.IndexHandle",
		otelTrace.WithSpanKind(otelTrace.SpanKindServer),
		otelTrace.WithNewRoot(),
		otelTrace.WithAttributes([]attribute.KeyValue{
			attribute.String("host", r.Host),
			attribute.String("path", r.URL.Path),
			attribute.String("method", r.Method),
		}...))
	defer span.End()

	view := &NotFoundView{
		Locale:          "fr",
		PageTitle:       "MedZoner.com - Not Found",
		TorHost:         r.Header.Get("TOR-HOST"),
		PageDescription: "MedZoner.com - Not Found",
	}
	_, err := h.Template.Render("404", view, w, http.StatusNotFound)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}
