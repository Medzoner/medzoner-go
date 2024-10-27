package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/router"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"github.com/dpapathanasiou/go-recaptcha"
)

type IServer interface {
	Start(ctx context.Context)
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
	conf config.Config,
	route router.IRouter,
	logger logger.ILogger,
	tracer tracer.Tracer,
) *Server {
	return &Server{
		Logger: logger,
		Router: route,
		HTTPServer: &http.Server{
			Addr:              fmt.Sprintf(":%d", conf.APIPort),
			Handler:           route,
			ReadHeaderTimeout: time.Second,
		},
		APIPort:            conf.APIPort,
		RecaptchaSecretKey: conf.RecaptchaSecretKey,
		Tracer:             tracer,
	}
}

func (s Server) Start(ctx context.Context) {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	recaptcha.Init(s.RecaptchaSecretKey)

	go func() {
		s.Logger.Log(fmt.Sprintf("Server up on port '%d'", s.APIPort))
		if err := s.HTTPServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.Logger.Error(fmt.Sprintf("listen and serve returned err: %v", err))
		}
	}()

	<-ctx.Done()
	s.Logger.Log("got interruption signal")
	if err := s.Shutdown(context.TODO()); err != nil { // Use here context with a required timeout
		s.Logger.Error(fmt.Sprintf("server shutdown returned an err: %v", err))
	}

	s.Logger.Log("server stopped!")
}

func (s Server) Shutdown(ctx context.Context) error {
	defer func() {
		if err := s.Tracer.ShutdownTracer(ctx); err != nil {
			s.Logger.Error(fmt.Sprintf("failed to shutdown TracerProvider: %s", err))
		}
		if err := s.Tracer.ShutdownMeter(ctx); err != nil {
			s.Logger.Error(fmt.Sprintf("failed to shutdown MeterProvider: %s", err))
		}
		if err := s.Tracer.ShutdownLogger(ctx); err != nil {
			s.Logger.Error(fmt.Sprintf("failed to shutdown LoggerProvider: %s", err))
		}
		s.Logger.Log("Tracer, Meter and Logger providers are shutdown")
	}()
	return s.HTTPServer.Shutdown(ctx)
}
