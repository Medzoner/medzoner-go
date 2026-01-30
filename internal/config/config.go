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
	"github.com/Medzoner/medzoner-go/pkg/captcha"
	"github.com/Medzoner/medzoner-go/pkg/database"
	"github.com/Medzoner/medzoner-go/pkg/notification"
)

type (
	RootPath string

	Config struct {
		Obs       *observability.Config `envPrefix:"TELEMETRY_"`
		Engine    ginadapter.Config     `envPrefix:"ENGINE_"`
		Logger    logger.Config         `envPrefix:"LOGGER_"`
		Auth      auth.Config
		Server    server.Config       `envPrefix:"SERVER_"`
		Mailer    notification.Config `envPrefix:"MAILER_"`
		Database  database.Config     `envPrefix:"DATABASE_"`
		RootPath  RootPath            `env:"ROOT_PATH"`
		Recaptcha captcha.Config      `envPrefix:"RECAPTCHA_"`
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
