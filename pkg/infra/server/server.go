package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/infra/router"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"github.com/dpapathanasiou/go-recaptcha"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
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
			Addr:    fmt.Sprintf(":%d", conf.APIPort),
			Handler: route,

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
			log.Fatalf("listen and serve returned err: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("got interruption signal")
	if err := s.Shutdown(context.TODO()); err != nil { // Use here context with a required timeout
		log.Printf("server shutdown returned an err: %v\n", err)
	}

	log.Println("server stopped!")
}

func (s Server) Shutdown(ctx context.Context) error {
	defer func() {
		if err := s.Tracer.ShutdownTracer(ctx); err != nil {
			log.Printf("failed to shutdown TracerProvider: %s", err)
		}
		if err := s.Tracer.ShutdownMeter(ctx); err != nil {
			log.Printf("failed to shutdown MeterProvider: %s", err)
		}
		if err := s.Tracer.ShutdownLogger(ctx); err != nil {
			log.Printf("failed to shutdown LoggerProvider: %s", err)
		}
	}()
	return s.HTTPServer.Shutdown(ctx)
}
