//+build wireinject

package main

import (
	"net/http"

	"github.com/cjmarkham/GoBB/src/repository"
	forumRepo "github.com/cjmarkham/GoBB/src/repository/forum"
	"github.com/google/wire"
	"github.com/gorilla/mux"

	"github.com/cjmarkham/GoBB/config"
	"github.com/cjmarkham/GoBB/src"
	"github.com/cjmarkham/GoBB/src/controllers"
	"github.com/cjmarkham/GoBB/src/controllers/home"
	"github.com/cjmarkham/GoBB/src/domain/forum"
)

var configProviders = wire.NewSet(
	config.ProvideServerConfig,
	config.ProvideLoggerConfig,
)

var controllerProviders = wire.NewSet(
	home.ProvideHomeController,
)

var serviceProviders = wire.NewSet(
	forum.ProvideService,
)

var repositoryProviders = wire.NewSet(
	config.ProvidePostgresqlConfig,
	repository.ProvideClient,
	forumRepo.ProvideRepository,
	wire.Bind(new(forum.Repository), new(*forumRepo.Repository)),
)

func InitializeApp() (*App, error) {
	panic(
		wire.Build(
			src.ProvideLogger,
			src.ProvideServer,
			ProvideApp,
			controller.ProvideRouter,
			wire.Bind(new(http.Handler), new(*mux.Router)),

			configProviders,
			controllerProviders,
			serviceProviders,
			repositoryProviders,
		),
	)
}
