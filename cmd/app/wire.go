//+build wireinject

package main

import (
	"net/http"

	"github.com/cjmarkham/GoBB/config"
	"github.com/google/wire"
	"github.com/gorilla/mux"

	"github.com/cjmarkham/GoBB"
	"github.com/cjmarkham/GoBB/controllers"
	"github.com/cjmarkham/GoBB/controllers/home"
)

var configProviders = wire.NewSet(
	config.ProvideServerConfig,
	config.ProvideLoggerConfig,
)

var controllerProviders = wire.NewSet(
	home.ProvideHomeController,
)

func InitializeApp() (*App, error) {
	panic(
		wire.Build(
			GoBB.ProvideLogger,
			GoBB.ProvideServer,
			ProvideApp,
			controllers.ProvideRouter,
			wire.Bind(new(http.Handler), new(*mux.Router)),

			configProviders,
			controllerProviders,
		),
	)
}
