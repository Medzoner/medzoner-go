package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/router"
)

// IServer Server Server
type IServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

// Server Server
type Server struct {
	HTTPServer *http.Server
}

// NewServer NewServer
func NewServer(conf config.IConfig, route router.IRouter) *Server {
	return &Server{
		HTTPServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", conf.GetAPIPort()),
			Handler: route,
		},
	}
}

func (s Server) ListenAndServe() error {
	return s.HTTPServer.ListenAndServe()
}

func (s Server) Shutdown(ctx context.Context) error {
	return s.HTTPServer.Shutdown(ctx)
}
