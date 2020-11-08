package server

import (
	"context"
	"net/http"
)

//Server Server
type IServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

//Server Server
type Server struct {
	HTTPServer *http.Server
}

func (s Server) ListenAndServe() error {
	return s.HTTPServer.ListenAndServe()
}

func (s Server) Shutdown(ctx context.Context) error {
	return s.HTTPServer.Shutdown(ctx)
}
