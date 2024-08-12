package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"

	env "github.com/caarlos0/env/v10"
)

// IConfig IConfig
type IConfig interface {
	Init() (*Config, error)
	GetRootPath() RootPath
	GetEnvironment() string
	GetMysqlDsn() string
	GetAPIPort() int
	GetDatabaseName() string
	GetDatabaseDriver() string
	GetMailerUser() string
	GetMailerPassword() string
	GetRecaptchaSiteKey() string
	GetRecaptchaSecretKey() string
	GetTraceFile() string
	GetOtelHost() string
}

type RootPath string

// Config Config
type Config struct {
	Environment        string `env:"ENV" envDefault:"dev"`
	RootPath           RootPath
	DebugMode          bool     `env:"DEBUG" envDefault:"false"`
	Options            []string `env:"OPTIONS" envDefault:"[]"`
	APIPort            int      `env:"API_PORT" envDefault:"8002"`
	DatabaseDsn        string   `env:"DATABASE_DSN" envDefault:"root:changeme@tcp(0.0.0.0:3366)"`
	DatabaseName       string   `env:"DATABASE_NAME" envDefault:"dev_medzoner"`
	DatabaseDriver     string   `env:"DATABASE_DRIVER" envDefault:"mysql"`
	MailerUser         string   `env:"MAILER_USER" envDefault:"medzoner@xxx.fake"`
	MailerPassword     string   `env:"MAILER_PASSWORD" envDefault:"xxxxxxxxxxxx"`
	RecaptchaSiteKey   string   `env:"RECAPTCHA_SITE_KEY" envDefault:"xxxxxxxxxxxx"`
	RecaptchaSecretKey string   `env:"RECAPTCHA_SECRET_KEY" envDefault:"xxxxxxxxxxxx"`
	TracerFile         string   `env:"TRACER_FILE" envDefault:"trace.out"`
}

// NewConfig NewConfig
func NewConfig() (*Config, error) {
	conf, err := parseEnv()
	if err != nil {
		return nil, err
	}
	return conf.Init()
}

// Init Init
func (c *Config) Init() (*Config, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	c.RootPath = RootPath(pwd + "/")
	err = godotenv.Load(string(c.RootPath) + "/.env")
	if err == nil {
		fmt.Println(".env file found")
	}
	if c.Environment == "test" {
		err = godotenv.Load(string(c.RootPath) + "/.env.test")
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

// GetMysqlDsn GetMysqlDsn
func (c *Config) GetMysqlDsn() string {
	return c.DatabaseDsn
}

// GetDatabaseDriver GetDatabaseDriver
func (c *Config) GetDatabaseDriver() string {
	return c.DatabaseDriver
}

// GetDatabaseName GetDatabaseName
func (c *Config) GetDatabaseName() string {
	return c.DatabaseName
}

// GetAPIPort GetAPIPort
func (c *Config) GetAPIPort() int {
	return c.APIPort
}

// GetRootPath GetRootPath
func (c *Config) GetRootPath() RootPath {
	return c.RootPath
}

// GetEnvironment GetEnvironment
func (c *Config) GetEnvironment() string {
	return c.Environment
}

// GetMailerUser GetMailerUser
func (c *Config) GetMailerUser() string {
	return c.MailerUser
}

// GetMailerPassword GetMailerPassword
func (c *Config) GetMailerPassword() string {
	return c.MailerPassword
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

// GetRecaptchaSiteKey GetRecaptchaSiteKey
func (c *Config) GetRecaptchaSiteKey() string {
	return c.RecaptchaSiteKey
}

// GetRecaptchaSecretKey GetRecaptchaSecretKey
func (c *Config) GetRecaptchaSecretKey() string {
	return c.RecaptchaSecretKey
}

// GetTraceFile GetTraceFile
func (c *Config) GetTraceFile() string {
	return c.TracerFile
}

// GetOtelHost GetOtelHost
func (c *Config) GetOtelHost() string {
	return getEnv("OTEL_HOST", "otel-collector-opentelemetry-collector.default.svc.cluster.local:4317")
}

// parseEnv parseEnv
func parseEnv() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
