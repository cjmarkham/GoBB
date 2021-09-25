package home

import (
	"html/template"
	"net/http"

	"github.com/rs/zerolog/log"
)

func (c Controller) Index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		log.Error().Err(err).Msg("could not parse templates")
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	if err := tpl.Execute(w, nil); err != nil {
		log.Error().Err(err).Msg("could not execute templates")
	}
}
