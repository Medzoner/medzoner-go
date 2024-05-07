package server

import (
	"context"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/router"
	"github.com/dpapathanasiou/go-recaptcha"
	"net/http"
	"time"
)

// IServer Server Server
type IServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

// Server Server
type Server struct {
	Logger             logger.ILogger
	Router             router.IRouter
	HTTPServer         *http.Server
	APIPort            int
	RecaptchaSecretKey string
}

// NewServer NewServer
func NewServer(
	conf config.IConfig,
	route router.IRouter,
	logger logger.ILogger,
) *Server {
	return &Server{
		Logger: logger,
		Router: route,
		HTTPServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", conf.GetAPIPort()),
			Handler: route,

			ReadHeaderTimeout: time.Second,
		},
		APIPort:            conf.GetAPIPort(),
		RecaptchaSecretKey: conf.GetRecaptchaSecretKey(),
	}
}

func (s Server) ListenAndServe() error {
	recaptcha.Init(s.RecaptchaSecretKey)

	s.Logger.Log(fmt.Sprintf("Server up on port '%d'", s.APIPort))
	return s.HTTPServer.ListenAndServe()
}

func (s Server) Shutdown(ctx context.Context) error {
	return s.HTTPServer.Shutdown(ctx)
}
