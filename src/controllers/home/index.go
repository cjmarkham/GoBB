package home

import (
	"context"
	"html/template"
	"net/http"

	"github.com/cjmarkham/GoBB/src/domain/forum"
	"github.com/rs/zerolog/log"
)

// TODO: Put somewhere else
type HomeIndexData struct {
	Title string
	Forums []forum.Forum
}

func (c Controller) Index(w http.ResponseWriter, r *http.Request) {
	forums, err := c.service.GetForums(context.TODO())
	if err != nil {
		log.Error().Err(err).Msg("could not get forums")
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	tpl, err := template.ParseFiles("src/views/home/index.html")
	if err != nil {
		log.Error().Err(err).Msg("could not parse templates")
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	data := HomeIndexData{
		Title: "Home",
		Forums: forums,
	}
	if err := tpl.Execute(w, data); err != nil {
		log.Error().Err(err).Msg("could not execute templates")
	}
}
