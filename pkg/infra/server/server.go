package server

import (
	"context"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/router"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"github.com/dpapathanasiou/go-recaptcha"
	"log"
	"net/http"
	"time"
)

type IServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type Server struct {
	Logger             logger.ILogger
	Router             router.IRouter
	HTTPServer         *http.Server
	APIPort            int
	RecaptchaSecretKey string
	Tracer             tracer.Tracer
}

// NewServer NewServer
func NewServer(
	conf config.IConfig,
	route router.IRouter,
	logger logger.ILogger,
	tracer tracer.Tracer,
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
		Tracer:             tracer,
	}
}

func (s Server) ListenAndServe() error {
	ctx := context.Background()
	defer func() {
		if err := s.Tracer.ShutdownTracer(ctx); err != nil {
			log.Printf("failed to shutdown TracerProvider: %s", err)
		}
	}()
	defer func() {
		if err := s.Tracer.ShutdownMeter(ctx); err != nil {
			log.Printf("failed to shutdown MeterProvider: %s", err)
		}
	}()
	recaptcha.Init(s.RecaptchaSecretKey)

	s.Logger.Log(fmt.Sprintf("Server up on port '%d'", s.APIPort))
	return s.HTTPServer.ListenAndServe()
}

func (s Server) Shutdown(ctx context.Context) error {
	return s.HTTPServer.Shutdown(ctx)
}
