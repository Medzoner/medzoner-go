package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"strconv"
	"strings"
)

type RootPath string

// Config Config
type Config struct {
	Environment        string `env:"ENV" envDefault:"dev"`
	RootPath           RootPath
	DebugMode          bool           `env:"DEBUG" envDefault:"false"`
	Options            []string       `env:"OPTIONS" envDefault:"[]"`
	APIPort            int            `env:"API_PORT" envDefault:"8002"`
	Database           DatabaseConfig `env:"DATABASE_"`
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
		return Config{}, err
	}
	cfg, err := conf.Init()
	if err != nil {
		return Config{}, err
	}
	return *cfg, nil
}

// Init Init
func (c *Config) Init() (*Config, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	c.RootPath = RootPath(pwd + "/")
	err = godotenv.Load(string(c.RootPath) + ".env")
	if err == nil {
		fmt.Println(".env file found")
	}
	if c.Environment == "test" {
		err = godotenv.Load(string(c.RootPath) + ".env.test")
		if err == nil {
			fmt.Println(".env.test file found")
		}
	}
	c.Options = getEnvAsSlice("OPTIONS", []string{}, ",")
	_ = getEnvAsBool("DEBUG_TEST", false)
	_ = getEnvAsInt("WAIT_MYSQL", 2)
	c.APIPort, _ = strconv.Atoi(getEnv("API_PORT", "8002"))

	return c, nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultVal
}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")
	if valStr == "" {
		return defaultVal
	}
	val := strings.Split(valStr, sep)
	return val
}

// GetOtelHost GetOtelHost
func (c *Config) GetOtelHost() string {
	//return getEnv("OTEL_HOST", "otel-collector-opentelemetry-collector.default.svc.cluster.local:4317")
	return getEnv("OTEL_HOST", "localhost:4317")
}

// parseEnv parseEnv
func parseEnv() (Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return Config{}, err
	}
	return *cfg, nil
}

// Debug Debug
func (c *Config) Debug() bool {
	return c.DebugMode
}
