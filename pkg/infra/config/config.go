package config

import (
	"os"

	"github.com/caarlos0/env/v10"
)

type RootPath string

// Config Config
type Config struct {
	Environment        string         `env:"ENV" envDefault:"dev"`
	RootPath           RootPath       `env:"ROOT_PATH"`
	DebugMode          bool           `env:"DEBUG" envDefault:"false"`
	Options            []string       `env:"OPTIONS" envDefault:"[]"`
	APIPort            int            `env:"API_PORT" envDefault:"8002"`
	Database           DatabaseConfig `envPrefix:"DATABASE_"`
	MailerUser         string         `env:"MAILER_USER" envDefault:"medzoner@xxx.fake"`
	MailerPassword     string         `env:"MAILER_PASSWORD" envDefault:"xxxxxxxxxxxx"`
	MailerHost         string         `env:"MAILER_HOST" envDefault:"smtp.gmail.com"`
	MailerPort         string         `env:"MAILER_PORT" envDefault:"587"`
	RecaptchaSiteKey   string         `env:"RECAPTCHA_SITE_KEY" envDefault:"xxxxxxxxxxxx"`
	RecaptchaSecretKey string         `env:"RECAPTCHA_SECRET_KEY" envDefault:"xxxxxxxxxxxx"`
	TracerFile         string         `env:"TRACER_FILE" envDefault:"trace.out"`
	OtelHost           string         `env:"OTEL_HOST" envDefault:"localhost:4317"`
}

// DatabaseConfig DatabaseConfig
type DatabaseConfig struct {
	Dsn    string `env:"DSN" envDefault:"root:changeme@tcp(0.0.0.0:3306)"`
	Name   string `env:"NAME" envDefault:"dev_medzoner"`
	Driver string `env:"DRIVER" envDefault:"mysql"`
}

// NewConfig is a constructor for Config
func NewConfig() (Config, error) {
	conf, err := parseEnv()
	if err != nil {
		return conf, err
	}
	if conf.RootPath == "" {
		pwd, err := os.Getwd()
		if err != nil {
			return conf, err
		}
		conf.RootPath = RootPath(pwd + "/")
	}

	return conf, nil
}

// parseEnv parseEnv
func parseEnv() (Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return *cfg, err
	}
	return *cfg, nil
}
