package config

import "github.com/caarlos0/env"

type PostgresqlConfig struct {
	DSN     string `env:"POSTGRES_DSN"`
	Timeout int    `env:"POSTGRES_TIMEOUT" envDefault:"5000"`
}

func ProvidePostgresqlConfig() (*PostgresqlConfig, error) {
	var config PostgresqlConfig
	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
