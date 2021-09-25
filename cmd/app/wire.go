//+build wireinject

package main

import (
	"net/http"

	"github.com/cjmarkham/GoBB/config"
	controllers2 "github.com/cjmarkham/GoBB/src/controllers"
	home2 "github.com/cjmarkham/GoBB/src/controllers/home"
	"github.com/google/wire"
	"github.com/gorilla/mux"

	"github.com/cjmarkham/GoBB"
)

var configProviders = wire.NewSet(
	config.ProvideServerConfig,
	config.ProvideLoggerConfig,
)

var controllerProviders = wire.NewSet(
	home2.ProvideHomeController,
)

func InitializeApp() (*App, error) {
	panic(
		wire.Build(
			GoBB.ProvideLogger,
			GoBB.ProvideServer,
			ProvideApp,
			controllers2.ProvideRouter,
			wire.Bind(new(http.Handler), new(*mux.Router)),

			configProviders,
			controllerProviders,
		),
	)
}
