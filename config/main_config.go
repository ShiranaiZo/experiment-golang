package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

var MainConfig AppConfig

type AppConfig struct {
	// Application
	AppName     string `env:"APP_NAME"`
	AppEnv      string `env:"APP_ENV"`
	AppKey      string `env:"APP_KEY"`
	AppHOST     string `env:"APP_HOST"`
	AppPORT     string `env:"APP_PORT"`
	AppTimezone string `env:"APP_TIMEZONE"`
	AppLocale   string `env:"APP_LOCALE"`

	// Api
	RouteAPIPath    string `env:"ROUTE_API_PATH"`
	RouteAPIVersion string `env:"ROUTE_API_VERSION"`

	// Cors
	CORSAllowedOrigins   string `env:"CORS_ALLOWED_ORIGINS"`
	CORSAllowedMethods   string `env:"CORS_ALLOWED_METHODS"`
	CORSAllowedHeaders   string `env:"CORS_ALLOWED_HEADERS"`
	CORSAllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS"`
	CORSExposeHeaders    string `env:"CORS_EXPOSE_HEADERS"`
	CORSMaxAge           int    `env:"CORS_MAX_AGE"`

	// Database
	DBConnection            string `env:"DB_CONNECTION"`
	DBHost                  string `env:"DB_HOST"`
	DBPort                  string `env:"DB_PORT"`
	DBName                  string `env:"DB_DATABASE"`
	DBUsername              string `env:"DB_USERNAME"`
	DBPassword              string `env:"DB_PASSWORD"`
	DBMaxOpenConnection     int    `env:"DB_MAX_OPEN_CONNECTION"`
	DBMaxLifetimeConnection int    `env:"DB_MAX_LIFETIME_CONNECTION"`
	DBMaxIdleConnection     int    `env:"DB_MAX_IDLE_CONNECTION"`
	DBMaxIdleTime           int    `env:"DB_MAX_IDLE_TIME"`
}

// GetEnv retrieves the value of the environment variable named by the key.
func InitMainConfig() {
	err := cleanenv.ReadEnv(&MainConfig)

	if err != nil {
		panic("Failed to read env on config: " + err.Error())
	}
}
