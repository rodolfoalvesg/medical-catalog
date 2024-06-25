package config

import (
	"strconv"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Loader is the interface that wraps the Load method.
type Loader interface {
	Load(cfg *Config) error
}

type Config struct {
	HTTP HTTPConfig
	DB   DBConfig
}

type HTTPConfig struct {
	Address      string        `envconfig:"HTTP_ADDRESS" default:":3000"`
	ReadTimeout  time.Duration `envconfig:"HTTP_READ_TIMEOUT" default:"5s"`
	WriteTimeout time.Duration `envconfig:"HTTP_WRITE_TIMEOUT" default:"10s"`
}

type DBConfig struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     int    `envconfig:"DB_PORT" default:"5432"`
	User     string `envconfig:"DB_USER" default:"postgres"`
	Password string `envconfig:"DB_PASSWORD" default:"postgres"`
	Database string `envconfig:"DB_DATABASE" default:"postgres"`
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

func (c *DBConfig) ConnectionString() string {
	port := strconv.FormatInt(int64(c.Port), 10)

	return "host=" + c.Host + " port=" + port + " user=" + c.User + " password=" + c.Password + " dbname=" + c.Database
}
