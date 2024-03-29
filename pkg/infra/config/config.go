package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
)

// IConfig IConfig
type IConfig interface {
	Init()
	GetRootPath() string
	GetEnvironment() string
	GetMysqlDsn() string
	GetAPIPort() int
	GetDatabaseName() string
	GetDatabaseDriver() string
	GetMailerUser() string
	GetMailerPassword() string
	GetRecaptchaSiteKey() string
	GetRecaptchaSecretKey() string
}

// Config Config
type Config struct {
	Environment        string
	RootPath           string
	DebugMode          bool
	Options            []string
	DatabaseDsn        string
	DatabaseName       string
	APIPort            int
	DatabaseDriver     string
	MailerUser         string
	MailerPassword     string
	RecaptchaSiteKey   string
	RecaptchaSecretKey string
}

// Init Init
func (c *Config) Init() {
	err := godotenv.Load(c.RootPath + "/.env")
	c.Environment = getEnv("ENV", "dev")
	if c.Environment == "test" {
		err = godotenv.Load(c.RootPath + "/.env.test")
	}
	c.MailerUser = getEnv("MAILER_USER", "medzoner@xxx.fake")
	c.MailerPassword = getEnv("MAILER_PASSWORD", "xxxxxxxxxxxx")
	c.DatabaseDsn = getEnv("DATABASE_DSN", "root:changeme@tcp(0.0.0.0:3366)")
	c.DatabaseDriver = getEnv("DATABASE_DRIVER", "mysql")
	c.DatabaseName = getEnv("DATABASE_NAME", "dev_medzoner")
	c.DebugMode = getEnvAsBool("DEBUG", false)
	c.Options = getEnvAsSlice("OPTIONS", []string{}, ",")
	_ = getEnvAsBool("DEBUG_TEST", false)
	_ = getEnvAsInt("WAIT_MYSQL", 2)
	c.RecaptchaSiteKey = getEnv("RECAPTCHA_SITE_KEY", "xxxxxxxxxxxx")
	c.RecaptchaSecretKey = getEnv("RECAPTCHA_SECRET_KEY", "xxxxxxxxxxxx")

	if err == nil {
		fmt.Println(".env file found")
		return
	}
	fmt.Println("No .env file found")
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
func (c *Config) GetRootPath() string {
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
