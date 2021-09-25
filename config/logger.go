package config

import "github.com/caarlos0/env"

type LoggerConfig struct {
	Level string `env:"LOG_LEVEL" envDefault:"debug"`
}

func ProvideLoggerConfig() (*LoggerConfig, error) {
	var config LoggerConfig
	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
