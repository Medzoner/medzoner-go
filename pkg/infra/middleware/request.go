package middleware

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

type correlationContextKey struct{}

// GetCorrelationID retrieves the correlation ID from the context.
func GetCorrelationID(ctx context.Context) string {
	if val := ctx.Value(correlationContextKey{}); val != nil {
		return val.(string)
	}
	return ""
}

// CorrelationMiddleware generates or propagates a correlation ID.
func (m APIMiddleware) CorrelationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = uuid.New().String()
		}

		ctx := context.WithValue(r.Context(), correlationContextKey{}, correlationID)
		w.Header().Set("X-Correlation-ID", correlationID)

		//ctx, cancel := context.WithTimeout(request.Context(), 60*time.Second)
		//defer cancel()
		//
		//ctx, span := h.Tracer.Start(
		//	ctx,
		//	"IndexHandler.IndexHandle",
		//	otelTrace.WithSpanKind(otelTrace.SpanKindServer),
		//	otelTrace.WithNewRoot(),
		//	otelTrace.WithAttributes([]attribute.KeyValue{
		//		attribute.String("host", request.Host),
		//		attribute.String("path", request.URL.Path),
		//		attribute.String("method", request.Method),
		//	}...))
		//defer span.End()

		lrw := NewLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r.WithContext(ctx))

		if lrw.statusCode == http.StatusInternalServerError {
			//resp := lrw.ResponseWriter.(http.ResponseWriter)
			body := "resp."
			m.Logger.Error("Internal server error : " + body)
			http.Error(lrw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			return
		}
	})
}

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}
