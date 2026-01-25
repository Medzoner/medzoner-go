package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/Medzoner/gomedz/pkg/observability"
	ginadapter "github.com/Medzoner/gomedz/pkg/http/adapter/gin"
	"github.com/Medzoner/gomedz/pkg/environment"
	"github.com/Medzoner/gomedz/pkg/auth"
	"github.com/Medzoner/gomedz/pkg/logger"
	"github.com/Medzoner/gomedz/pkg/http/server"
	"github.com/Medzoner/gomedz/pkg/config"
)

type Config struct {
	Obs                *observability.Config `envPrefix:"TELEMETRY_"`
	Engine             ginadapter.Config     `envPrefix:"ENGINE_"`
	AppName            string                `env:"APP_NAME"         envDefault:"audio-remover"`
	Env                environment.Env       `env:"ENV"              envDefault:"development"`
	Logger             logger.Config         `envPrefix:"LOGGER_"`
	Auth               auth.Config
	Server             server.Config  `envPrefix:"SERVER_"`
	Mailer             MailerConfig   `envPrefix:"MAILER_"`
	Database           DatabaseConfig `envPrefix:"DATABASE_"`
	RootPath           RootPath       `env:"ROOT_PATH"`
	RecaptchaSiteKey   string         `env:"RECAPTCHA_SITE_KEY"   envDefault:"xxxxxxxxxxxx"`
	RecaptchaSecretKey string         `env:"RECAPTCHA_SECRET_KEY" envDefault:"xxxxxxxxxxxx"`
}

func NewConfig() (Config, error) {
	cfg := Config{}

	if err := env.ParseWithOptions(&cfg, env.Options{}); err != nil {
		return cfg, fmt.Errorf("failed to parse environment variables: %w", err)
	}

	config.PrintConfigEnv(cfg, "")

	return cfg, nil
}

type RootPath string

type MailerConfig struct {
	User     string `env:"USER"     envDefault:"medzoner@xxx.fake"`
	Password string `env:"PASSWORD" envDefault:"xxxxxxxxxxxx"`
	Host     string `env:"HOST"     envDefault:"smtp.gmail.com"`
	Port     string `env:"PORT"     envDefault:"587"`
}

type DatabaseConfig struct {
	Dsn    string `env:"DSN"    envDefault:"root:changeme@tcp(0.0.0.0:3306)"`
	Name   string `env:"NAME"   envDefault:"dev_medzoner"`
	Driver string `env:"DRIVER" envDefault:"mysql"`
}
