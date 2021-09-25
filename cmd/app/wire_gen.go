// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/cjmarkham/GoBB"
	"github.com/cjmarkham/GoBB/config"
	"github.com/cjmarkham/GoBB/controllers"
	"github.com/cjmarkham/GoBB/controllers/home"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeApp() (*App, error) {
	controller, err := home.ProvideHomeController()
	if err != nil {
		return nil, err
	}
	loggerConfig, err := config.ProvideLoggerConfig()
	if err != nil {
		return nil, err
	}
	logger, err := GoBB.ProvideLogger(loggerConfig)
	if err != nil {
		return nil, err
	}
	router := controllers.ProvideRouter(controller, logger)
	serverConfig, err := config.ProvideServerConfig()
	if err != nil {
		return nil, err
	}
	server := GoBB.ProvideServer(router, serverConfig)
	app := ProvideApp(server)
	return app, nil
}

// wire.go:

var configProviders = wire.NewSet(config.ProvideServerConfig, config.ProvideLoggerConfig)

var controllerProviders = wire.NewSet(home.ProvideHomeController)
