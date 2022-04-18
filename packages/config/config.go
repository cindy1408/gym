package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseURL       string `envconfig:"DATABASE_URL" default:"postgresql://user:password@localhost:5432/gym?sslmode=disable"`
	Host              string `envconfig:"HOST" default:"0.0.0.0"`
	Port              string `envconfig:"PORT" default:"8080"`
	GraphQLPlayground bool   `envconfig:"GRAPHQL_PLAYGROUND" default:"false"`
	DisableDatabase   bool   `envconfig:"DISABLE_DATABASE" default:"false"`
}

func Load() (*Config, error) {
	var c Config
	err := envconfig.Process("cal", &c)
	return &c, err
}

// MustLoad will Load all config vars or cause a fatal exit
func MustLoadConfig() *Config {
	config, err := Load()
	if err != nil {
		log.Fatalf("failed to load env config: %v", err)
	}
	return config
}
