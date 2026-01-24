package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/Medzoner/gomedz/pkg/observability"
	ginadapter "github.com/Medzoner/gomedz/pkg/http/adapter/gin"
	"github.com/Medzoner/gomedz/pkg/environment"
	"github.com/Medzoner/gomedz/pkg/auth"
	"github.com/Medzoner/gomedz/pkg/logger"
	"github.com/Medzoner/gomedz/pkg/http/server"
	"github.com/Medzoner/gomedz/pkg/config"
)

type Config2 struct {
	Obs     observability.Config `envPrefix:"TELEMETRY_"`
	Engine  ginadapter.Config    `envPrefix:"ENGINE_"`
	AppName string               `env:"APP_NAME"         envDefault:"audio-remover"`
	Env     environment.Env      `env:"ENV"              envDefault:"development"`
	Logger  logger.Config        `envPrefix:"LOGGER_"`
	Auth    auth.Config
	Server  server.Config `envPrefix:"SERVER_"`
}

func NewConfig2() (*Config2, error) {
	cfg := &Config2{}

	if err := env.ParseWithOptions(cfg, env.Options{}); err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %w", err)
	}

	config.PrintConfigEnv(*cfg, "")

	return cfg, nil
}

type RootPath string

type Config struct {
	Mailer             MailerConfig   `envPrefix:"MAILER_"`
	Database           DatabaseConfig `envPrefix:"DATABASE_"`
	Environment        string         `env:"ENV"                  envDefault:"dev"`
	RootPath           RootPath       `env:"ROOT_PATH"`
	RecaptchaSiteKey   string         `env:"RECAPTCHA_SITE_KEY"   envDefault:"xxxxxxxxxxxx"`
	RecaptchaSecretKey string         `env:"RECAPTCHA_SECRET_KEY" envDefault:"xxxxxxxxxxxx"`
	OtelHost           string         `env:"OTEL_HOST"            envDefault:"localhost:4317"`
	Options            []string       `env:"OPTIONS"              envDefault:"[]"`
	APIPort            int            `env:"API_PORT"             envDefault:"8002"`
	DebugMode          bool           `env:"DEBUG"                envDefault:"false"`
}

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

// NewConfig is a constructor for Config
func NewConfig() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return cfg, fmt.Errorf("parse env: %w", err)
	}
	if cfg.RootPath == "" {
		pwd, err := os.Getwd()
		if err != nil {
			return cfg, fmt.Errorf("get current working directory: %w", err)
		}
		cfg.RootPath = RootPath(pwd + "/")
	}

	return cfg, nil
}
