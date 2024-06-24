package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Loader is the interface that wraps the Load method.
type Loader interface {
	Load(cfg *Config) error
}

type Config struct {
	HTTP HTTPConfig
}

type HTTPConfig struct {
	Address      string        `envconfig:"HTTP_ADDRESS" default:":3000"`
	ReadTimeout  time.Duration `envconfig:"HTTP_READ_TIMEOUT" default:"5s"`
	WriteTimeout time.Duration `envconfig:"HTTP_WRITE_TIMEOUT" default:"10s"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	noPrefix := ""

	// Load the configuration from the environment
	if err := envconfig.Process(noPrefix, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
