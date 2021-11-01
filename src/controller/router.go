package controller

import (
	"net/http"

	"github.com/cjmarkham/GoBB/src/controller/topic"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	"github.com/cjmarkham/GoBB/src/controller/forum"
	"github.com/cjmarkham/GoBB/src/controller/home"
)

func ProvideRouter(
	homeController *home.Controller,
	forumController *forum.Controller,
	topicController *topic.Controller,
	logger zerolog.Logger,
) *mux.Router {
	router := &mux.Router{}

	router.Handle("/f/{forumSlug}", http.HandlerFunc(forumController.Show)).Methods("GET")
	router.Handle("/f/{forumSlug}/t/{topicSlug}", http.HandlerFunc(topicController.Show)).Methods("GET")
	router.Handle("/", http.HandlerFunc(homeController.Index)).Methods("GET")

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	return router
}
