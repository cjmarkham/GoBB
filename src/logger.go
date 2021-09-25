package src

import (
	"fmt"
	"os"

	"github.com/cjmarkham/GoBB/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ProvideLogger(config *config.LoggerConfig) (zerolog.Logger, error) {
	level, err := zerolog.ParseLevel(config.Level)
	if err != nil {
		return zerolog.Logger{}, fmt.Errorf("invalid log level %s", config.Level)
	}
	zerolog.SetGlobalLevel(level)
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
		With().
		Timestamp().
		Logger()

	log.Logger = logger

	return logger, nil
}
