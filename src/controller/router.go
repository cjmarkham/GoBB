package controller

import (
	"net/http"

	home2 "github.com/cjmarkham/GoBB/src/controllers/home"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

func ProvideRouter(
	homeController *home2.Controller,
	logger zerolog.Logger,
) *mux.Router {
	router := &mux.Router{}

	router.Handle("/", http.HandlerFunc(homeController.Index)).Methods("GET")

	return router
}
