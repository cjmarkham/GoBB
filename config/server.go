package config

import "github.com/caarlos0/env"

type ServerConfig struct {
	Port int `env:"PORT" envDefault:"8000"`
}

func ProvideServerConfig() (*ServerConfig, error) {
	var config ServerConfig
	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
