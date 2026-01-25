package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"

	ginadapter "github.com/Medzoner/gomedz/pkg/http/adapter/gin"

	"github.com/Medzoner/gomedz/pkg/auth"
	"github.com/Medzoner/gomedz/pkg/logger"
	"github.com/Medzoner/gomedz/pkg/http/server"
	"github.com/Medzoner/gomedz/pkg/config"
	"github.com/Medzoner/gomedz/pkg/observability"
)

type (
	RootPath string

	Config struct {
		Obs       *observability.Config `envPrefix:"TELEMETRY_"`
		Engine    ginadapter.Config     `envPrefix:"ENGINE_"`
		Logger    logger.Config         `envPrefix:"LOGGER_"`
		Auth      auth.Config
		Server    server.Config `envPrefix:"SERVER_"`
		Mailer    Mailer        `envPrefix:"MAILER_"`
		Database  Database      `envPrefix:"DATABASE_"`
		RootPath  RootPath      `env:"ROOT_PATH"`
		Recaptcha Recaptcha     `envPrefix:"RECAPTCHA_"`
	}

	Mailer struct {
		User     string `env:"USER"     envDefault:"medzoner@xxx.fake"`
		Password string `env:"PASSWORD" envDefault:"xxxxxxxxxxxx"`
		Host     string `env:"HOST"     envDefault:"smtp.gmail.com"`
		Port     string `env:"PORT"     envDefault:"587"`
	}

	Recaptcha struct {
		RecaptchaSiteKey   string `env:"SITE_KEY"   envDefault:"xxxxxxxxxxxx"`
		RecaptchaSecretKey string `env:"SECRET_KEY" envDefault:"xxxxxxxxxxxx"`
	}

	Database struct {
		Dsn    string `env:"DSN"    envDefault:"root:changeme@tcp(0.0.0.0:3306)"`
		Name   string `env:"NAME"   envDefault:"dev_medzoner"`
		Driver string `env:"DRIVER" envDefault:"mysql"`
	}
)

func NewConfig() (Config, error) {
	cfg := Config{}

	if err := env.ParseWithOptions(&cfg, env.Options{}); err != nil {
		return cfg, fmt.Errorf("failed to parse environment variables: %w", err)
	}

	config.PrintConfigEnv(cfg, "")

	return cfg, nil
}
