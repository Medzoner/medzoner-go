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
	statusCode int
	body       []byte
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
	return lrw.ResponseWriter.Write(body)
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK, make([]byte, 0)}
}
