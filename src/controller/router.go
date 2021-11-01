package controller

import (
	"net/http"

	"github.com/cjmarkham/GoBB/src/controller/forum"
	"github.com/cjmarkham/GoBB/src/controller/home"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

func ProvideRouter(
	homeController *home.Controller,
	forumController *forum.Controller,
	logger zerolog.Logger,
) *mux.Router {
	router := &mux.Router{}

	router.Handle("/f/{slug}", http.HandlerFunc(forumController.Index)).Methods("GET")
	router.Handle("/", http.HandlerFunc(homeController.Index)).Methods("GET")

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	return router
}
