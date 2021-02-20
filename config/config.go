package config

import (
	"go-echo-template/logger"
	"time"

	"github.com/kelseyhightower/envconfig"
)

var (
	// Cfg app global
	Cfg *Config
)

// Config structure
type Config struct {
	Version       string        `envconfig:"APP_VERSION" default:"0.0.0"`
	Revision      string        `envconfig:"APP_REVISION" default:""`
	ServerHost    string        `envconfig:"APP_SERVER_HOST" default:"0.0.0.0"`
	ServerPort    string        `envconfig:"APP_SERVER_PORT" default:"8000"`
	ReadTimeout   time.Duration `envconfig:"APP_SERVER_READ_TIMEOUT" default:"60s"`
	WriteTimeout  time.Duration `envconfig:"APP_SERVER_WRITE_TIMEOUT" default:"60s"`
	SentryDSN     string        `envconfig:"APP_SENTRY_DSN" default:""`
	DBConnStr     string        `envconfig:"APP_DATABASE_DSN"`
	RunMigrations bool          `envconfig:"APP_DATABASE_RUN_MIGRATIONS" default:"true"`
}

// Get config
func Get() *Config {
	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		logger.Log.Fatal(err)
	}

	return cfg
}
