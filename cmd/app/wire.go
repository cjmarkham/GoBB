//+build wireinject

package main

import (
	"net/http"

	"github.com/google/wire"
	"github.com/gorilla/mux"

	"github.com/cjmarkham/GoBB/config"
	"github.com/cjmarkham/GoBB/src"
	"github.com/cjmarkham/GoBB/src/controller"
	forumController "github.com/cjmarkham/GoBB/src/controller/forum"
	homeController "github.com/cjmarkham/GoBB/src/controller/home"
	topicController "github.com/cjmarkham/GoBB/src/controller/topic"
	"github.com/cjmarkham/GoBB/src/domain/forum"
	"github.com/cjmarkham/GoBB/src/domain/post"
	"github.com/cjmarkham/GoBB/src/domain/topic"
	"github.com/cjmarkham/GoBB/src/domain/user"
	"github.com/cjmarkham/GoBB/src/repository"
	forumRepo "github.com/cjmarkham/GoBB/src/repository/forum"
	topicRepo "github.com/cjmarkham/GoBB/src/repository/topic"
	postRepo "github.com/cjmarkham/GoBB/src/repository/post"
)

var configProviders = wire.NewSet(
	config.ProvideServerConfig,
	config.ProvideLoggerConfig,
)

var controllerProviders = wire.NewSet(
	homeController.ProvideController,
	forumController.ProvideController,
	topicController.ProvideController,
)

var serviceProviders = wire.NewSet(
	forum.ProvideService,
	topic.ProvideService,
	post.ProvideService,
	user.ProvideService,
)

var repositoryProviders = wire.NewSet(
	config.ProvidePostgresqlConfig,
	repository.ProvideClient,
	forumRepo.ProvideRepository,
	wire.Bind(new(forum.Repository), new(*forumRepo.Repository)),
	topicRepo.ProvideRepository,
	wire.Bind(new(topic.Repository), new(*topicRepo.Repository)),
	postRepo.ProvideRepository,
	wire.Bind(new(post.Repository), new(*postRepo.Repository)),
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
