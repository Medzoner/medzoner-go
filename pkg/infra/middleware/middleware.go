package middleware

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
)

// APIMiddleware APIMiddleware
type APIMiddleware struct {
	Logger logger.ILogger
}

// NewAPIMiddleware is a factory function to create a new APIMiddleware
func NewAPIMiddleware(logger logger.ILogger) APIMiddleware {
	return APIMiddleware{
		Logger: logger,
	}
}
