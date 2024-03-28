package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	UserServiceHost string
	UserServicePort int

	ProductServiceHost string
	ProductServicePort int

	AccessTokenTimout int // MINUTES
	SignInKey         string

	AuthConfigPath string
	CSVFilePath    string

	// context timeout in seconds
	CtxTimeout int

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))
	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "localhost"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 9000))

	c.ProductServiceHost = cast.ToString(getOrReturnDefault("PRODUCT_SERVICE_HOST", "localhost"))
	c.ProductServicePort = cast.ToInt(getOrReturnDefault("PRODUCT_SERVICE_PORT", 1212))

	c.AccessTokenTimout = cast.ToInt(getOrReturnDefault("ACCESS_TOKEN_TIMEOUT", 300))
	c.SignInKey = cast.ToString(getOrReturnDefault("SIGN_IN_KEY", "test_key"))

	// casbin...
	c.AuthConfigPath = cast.ToString(getOrReturnDefault("AUTH_PATH", "./config/auth.conf"))
	c.CSVFilePath = cast.ToString(getOrReturnDefault("CSV_FILE_PATH", "./config/auth.csv"))

	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
