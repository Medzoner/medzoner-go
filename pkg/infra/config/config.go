package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
)

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
}

type Config struct {
	Environment    string
	RootPath       string
	DebugMode      bool
	Options        []string
	DatabaseDsn    string
	DatabaseName   string
	ApiPort        int
	DatabaseDriver string
	MailerUser     string
	MailerPassword string
}

func (c *Config) Init() {
	err := godotenv.Load(c.RootPath + "/.env")
	c.Environment = getEnv("ENV", "test")
	if c.Environment == "test" {
		err = godotenv.Load(c.RootPath + "/.env.test")
	}
	c.MailerUser = getEnv("MAILER_USER", "medzoner@xxx.fake")
	c.MailerPassword = getEnv("MAILER_PASSWORD", "xxxxxxxxxxxx")
	c.DatabaseDsn = getEnv("DATABASE_DSN", "dev")
	c.DatabaseDriver = getEnv("DATABASE_DRIVER", "mysql")
	c.DatabaseName = getEnv("DATABASE_NAME", "dev")
	c.DebugMode = getEnvAsBool("DEBUG", false)
	c.Options = getEnvAsSlice("OPTIONS", []string{}, ",")
	_ = getEnvAsBool("DEBUG_TEST", false)
	_ = getEnvAsInt("WAIT_MYSQL", 2)

	if err == nil {
		fmt.Println(".env file found")
		return
	}
	fmt.Println("No .env file found")
}

func (c *Config) GetMysqlDsn() string {
	return c.DatabaseDsn
}

func (c *Config) GetDatabaseDriver() string {
	return c.DatabaseDriver
}

func (c *Config) GetDatabaseName() string {
	return c.DatabaseName
}

func (c *Config) GetAPIPort() int {
	return c.ApiPort
}

func (c *Config) GetRootPath() string {
	return c.RootPath
}

func (c *Config) GetEnvironment() string {
	return c.Environment
}

func (c *Config) GetMailerUser() string {
	return c.MailerUser
}

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
