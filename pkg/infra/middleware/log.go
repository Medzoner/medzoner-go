package middleware

import (
	"fmt"
	"net/http"
)

// LogMiddleware is a middleware that logs the request and response
func (m APIMiddleware) LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := NewLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r.WithContext(r.Context()))

		if lrw.statusCode == http.StatusInternalServerError {
			m.Telemetry.Error(r.Context(), fmt.Sprintf("Internal server error : %s", string(lrw.body)))

			http.Error(lrw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	})
}

type LoggingResponseWriter struct {
	http.ResponseWriter
	body       []byte
	statusCode int
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *LoggingResponseWriter) Write(body []byte) (int, error) {
	if lrw.statusCode == http.StatusInternalServerError && len(lrw.body) == 0 {
		lrw.body = body
		return 1, nil
	}
	i, err := lrw.ResponseWriter.Write(body)
	if err != nil {
		return 0, fmt.Errorf("error writing response: %w", err)
	}
	return i, nil
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, make([]byte, 0), http.StatusOK}
}
