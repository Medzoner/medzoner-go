package middleware

import (
	"net/http"
)

// IMiddleware IMiddleware
type IMiddleware interface {
	Middleware(next http.Handler) http.Handler
}

// APIMiddleware APIMiddleware
type APIMiddleware struct{}

// NewAPIMiddleware NewAPIMiddleware
func NewAPIMiddleware() IMiddleware {
	return &APIMiddleware{}

}

// Middleware Middleware
func (m APIMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")

		next.ServeHTTP(w, r)
	})
}
