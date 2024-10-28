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
	ShutdownWithTimeout() error
}

type Server struct {
	Logger             logger.ILogger
	Router             router.IRouter
	HTTPServer         *http.Server
	APIPort            int
	RecaptchaSecretKey string
	Tracer             tracer.Tracer
}

// NewServer initializes a new Server instance with configurations.
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

// Start initiates the server and listens for system signals for graceful shutdown.
func (s *Server) Start(ctx context.Context) {
	// Capture OS signals for graceful shutdown
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	recaptcha.Init(s.RecaptchaSecretKey)
	s.Logger.Log(fmt.Sprintf("Server starting on port '%d'", s.APIPort))

	// Run server in a goroutine to handle requests
	go func() {
		if err := s.HTTPServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.Logger.Error(fmt.Sprintf("HTTP server error: %v", err))
		}
	}()

	<-ctx.Done() // Wait for signal

	s.Logger.Log("Interrupt signal received; shutting down server...")
	if err := s.ShutdownWithTimeout(); err != nil {
		s.Logger.Error(fmt.Sprintf("Shutdown error: %v", err))
	}

	s.Logger.Log("Server stopped successfully!")
}

// ShutdownWithTimeout gracefully shuts down the server with a timeout context.
func (s *Server) ShutdownWithTimeout() error {
	// Set a timeout for the shutdown process
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown HTTP server and then tracer, meter, and logger providers
	if err := s.HTTPServer.Shutdown(ctx); err != nil {
		s.Logger.Error(fmt.Sprintf("HTTP server shutdown error: %v", err))
	}

	// Shutdown tracer components
	if err := s.shutdownTracerProviders(ctx); err != nil {
		return fmt.Errorf("tracer providers shutdown error: %w", err)
	}

	return nil
}

// shutdownTracerProviders shuts down all tracer-related providers in the context.
func (s *Server) shutdownTracerProviders(ctx context.Context) error {
	if err := s.Tracer.ShutdownTracer(ctx); err != nil {
		s.logProviderShutdownError("TracerProvider", err)
	}
	if err := s.Tracer.ShutdownMeter(ctx); err != nil {
		s.logProviderShutdownError("MeterProvider", err)
	}
	if err := s.Tracer.ShutdownLogger(ctx); err != nil {
		s.logProviderShutdownError("LoggerProvider", err)
	}
	s.Logger.Log("Tracer, Meter, and Logger providers have been shut down")
	return nil
}

// logProviderShutdownError logs errors for shutdown of specific providers.
func (s *Server) logProviderShutdownError(provider string, err error) {
	if err != nil {
		s.Logger.Error(fmt.Sprintf("Failed to shutdown %s: %v", provider, err))
	}
}
