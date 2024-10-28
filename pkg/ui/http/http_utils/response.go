package http_utils

import (
	"fmt"
	"net/http"

	otelTrace "go.opentelemetry.io/otel/trace"
)

func ResponseError(w http.ResponseWriter, err error, code int, span otelTrace.Span) {
	span.RecordError(err)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	i, err := w.Write([]byte(err.Error()))
	if err != nil {
		fmt.Println(i, err)
	}
}
