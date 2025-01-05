package middleware

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/telemetry"
)

// APIMiddleware APIMiddleware
type APIMiddleware struct {
	Telemetry telemetry.Telemeter
}

// NewAPIMiddleware is a factory function to create a new APIMiddleware
func NewAPIMiddleware(tm telemetry.Telemeter) APIMiddleware {
	return APIMiddleware{
		Telemetry: tm,
	}
}
