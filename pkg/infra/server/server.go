package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/router"
	"github.com/Medzoner/medzoner-go/pkg/infra/telemetry"
	"github.com/dpapathanasiou/go-recaptcha"
)

type IServer interface {
	Start(ctx context.Context)
	Shutdown(ctx context.Context) error
}

type Server struct {
	Router             router.IRouter
	Telemetry          telemetry.Telemeter
	HTTPServer         *http.Server
	RecaptchaSecretKey string
	APIPort            int
}

// NewServer initializes a new Server instance with configurations.
func NewServer(
	conf config.Config,
	route router.IRouter,
	tm telemetry.Telemeter,
) *Server {
	return &Server{
		Router: route,
		HTTPServer: &http.Server{
			Addr:              fmt.Sprintf(":%d", conf.APIPort),
			Handler:           route,
			ReadHeaderTimeout: time.Second,
		},
		APIPort:            conf.APIPort,
		RecaptchaSecretKey: conf.RecaptchaSecretKey,
		Telemetry:          tm,
	}
}

// Start initiates the server and listens for system signals for graceful shutdown.
func (s *Server) Start(ctx context.Context) {
	s.profile(ctx)

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	recaptcha.Init(s.RecaptchaSecretKey)
	s.Telemetry.Log(ctx, fmt.Sprintf("Server starting on port :%d", s.APIPort))

	go func() {
		if err := s.HTTPServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.Telemetry.Error(ctx, fmt.Sprintf("HTTP server error: %v", err))
		}
	}()

	<-ctx.Done()

	s.Telemetry.Log(ctx, "Interrupt signal received; shutting down server...")
	if err := s.Shutdown(ctx); err != nil {
		s.Telemetry.Error(ctx, fmt.Sprintf("Shutdown error: %v", err))
	}

	s.Telemetry.Log(ctx, "Server stopped successfully!")
}

func (s *Server) profile(ctx context.Context) {
	if os.Getenv("DEBUG") == "true" {
		s.Telemetry.Log(ctx, "Starting pprof server on :6060")
		go func() {
			s.Telemetry.Log(ctx, fmt.Sprintf("error : %f", http.ListenAndServe("localhost:6060", nil)))
		}()
	}
}

// Shutdown gracefully shuts down the server with a timeout context.
func (s *Server) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := s.HTTPServer.Shutdown(ctx); err != nil {
		s.Telemetry.Error(ctx, fmt.Sprintf("HTTP server shutdown error: %v", err))
	}

	if err := s.shutdownTracerProviders(ctx); err != nil {
		return fmt.Errorf("telemetry providers shutdown error: %w", err)
	}

	return nil
}

// shutdownTracerProviders shuts down all telemetry-related providers in the context.
func (s *Server) shutdownTracerProviders(ctx context.Context) error {
	if err := s.Telemetry.ShutdownTracer(ctx); err != nil {
		s.logProviderShutdownError(ctx, "TracerProvider", err)
	}
	if err := s.Telemetry.ShutdownMeter(ctx); err != nil {
		s.logProviderShutdownError(ctx, "MeterProvider", err)
	}
	if err := s.Telemetry.ShutdownLogger(ctx); err != nil {
		s.logProviderShutdownError(ctx, "LoggerProvider", err)
	}
	s.Telemetry.Log(ctx, "Telemeter, Meter, and Telemetry providers have been shut down")
	return nil
}

// logProviderShutdownError logs errors for shutdown of specific providers.
func (s *Server) logProviderShutdownError(ctx context.Context, provider string, err error) {
	if err != nil {
		s.Telemetry.Error(ctx, fmt.Sprintf("Failed to shutdown %s: %v", provider, err))
	}
}
