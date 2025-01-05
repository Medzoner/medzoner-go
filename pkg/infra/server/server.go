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
	"github.com/Medzoner/medzoner-go/pkg/infra/router"
	"github.com/Medzoner/medzoner-go/pkg/infra/telemetry"
	"github.com/dpapathanasiou/go-recaptcha"
)

type IServer interface {
	Start(ctx context.Context)
	ShutdownWithTimeout() error
}

type Server struct {
	Router             router.IRouter
	HTTPServer         *http.Server
	APIPort            int
	RecaptchaSecretKey string
	Telemetry          telemetry.Telemeter
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
	// Capture OS signals for graceful shutdown
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	recaptcha.Init(s.RecaptchaSecretKey)
	s.Telemetry.Log(ctx, fmt.Sprintf("Server starting on port '%d'", s.APIPort))

	// Run server in a goroutine to handle requests
	go func() {
		if err := s.HTTPServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.Telemetry.Error(ctx, fmt.Sprintf("HTTP server error: %v", err))
		}
	}()

	<-ctx.Done() // Wait for signal

	s.Telemetry.Log(ctx, "Interrupt signal received; shutting down server...")
	if err := s.ShutdownWithTimeout(); err != nil {
		s.Telemetry.Error(ctx, fmt.Sprintf("Shutdown error: %v", err))
	}

	s.Telemetry.Log(ctx, "Server stopped successfully!")
}

// ShutdownWithTimeout gracefully shuts down the server with a timeout context.
func (s *Server) ShutdownWithTimeout() error {
	// Set a timeout for the shutdown process
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown HTTP server and then telemetry, meter, and logger providers
	if err := s.HTTPServer.Shutdown(ctx); err != nil {
		s.Telemetry.Error(ctx, fmt.Sprintf("HTTP server shutdown error: %v", err))
	}

	// Shutdown telemetry components
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
