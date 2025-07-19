package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	APIKey  string `default:"11111111-1111-1111-1111-111111111111" env:"API_KEY"`
	Referer string `default:"https://docs.rarible.org" env:"REFERER"`
}

func ReadConfig() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}
