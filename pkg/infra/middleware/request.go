package middleware

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

type CorrelationContextKey struct{}

// GetCorrelationID retrieves the correlation ID from the context.
func GetCorrelationID(ctx context.Context) string {
	if val := ctx.Value(CorrelationContextKey{}); val != nil {
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
		ctx := context.WithValue(r.Context(), CorrelationContextKey{}, correlationID)
		w.Header().Set("X-Correlation-ID", correlationID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
