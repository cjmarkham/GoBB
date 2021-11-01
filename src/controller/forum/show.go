package forum

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/cjmarkham/GoBB/src/controller/helpers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func (c Controller) Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	forum, err := c.forumService.GetBySlug(context.TODO(), params["forumSlug"])
	if err != nil {
		log.Error().Err(err).Msg("could not get forum")
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	topics, err := c.topicService.GetForumTopics(context.TODO(), forum.ID)
	if err != nil {
		log.Error().Err(err).Msg("could not get forum topics")
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	tpl, err := template.ParseFiles(
		"src/views/layout.html",
		"src/views/header.html",
		"src/views/forum/show.html",
	)

	if err != nil {
		log.Error().Err(err).Msg("could not parse templates")
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	data := IndexData{
		Title: forum.Name,
		Forum: forum,
		Topics: topics,
		Crumbs: []helpers.Crumb{
			{
				Name: "Home",
				Link: "/",
			},
			{
				Name: forum.Name,
				Link: fmt.Sprintf("/f/%s", forum.Slug),
			},
		},
	}
	if err := tpl.Execute(w, data); err != nil {
		log.Error().Err(err).Msg("could not execute templates")
	}
}
