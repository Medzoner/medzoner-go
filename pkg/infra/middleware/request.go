package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
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
func CorrelationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = uuid.New().String()
		}

		ctx := context.WithValue(r.Context(), correlationContextKey{}, correlationID)

		w.Header().Set("X-Correlation-ID", correlationID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
