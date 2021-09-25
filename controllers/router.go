package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	"github.com/cjmarkham/GoBB/controllers/home"
)

func ProvideRouter(
	homeController *home.Controller,
	logger zerolog.Logger,
) *mux.Router {
	router := &mux.Router{}

	router.Handle("/", http.HandlerFunc(homeController.Index)).Methods("GET")

	return router
}
