package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
)

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
