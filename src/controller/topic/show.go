package topic

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

	topic, err := c.topicService.GetBySlug(context.TODO(), params["topicSlug"])
	if err != nil {
		log.Error().Err(err).Msg("could not get topic")
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	posts, err := c.postService.GetTopicPosts(context.TODO(), topic.ID)
	if err != nil {
		log.Error().Err(err).Msg("could not get topic posts")
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	tpl, err := template.ParseFiles(templates()...)

	if err != nil {
		log.Error().Err(err).Msg("could not parse templates")
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	data := IndexData{
		Title: topic.Name,
		Posts: posts,
		Crumbs: []helpers.Crumb{
			{
				Name: "Home",
				Link: "/",
			},
			{
				Name: forum.Name,
				Link: fmt.Sprintf("/f/%s", forum.Slug),
			},
			{
				Name: topic.Name,
				Link: fmt.Sprintf("/f/%s/t/%s", forum.Slug, topic.Slug),
			},
		},
	}
	if err := tpl.Execute(w, data); err != nil {
		log.Error().Err(err).Msg("could not execute templates")
	}
}

func templates() []string {
	templates := helpers.GlobalTemplates()
	templates = append(templates,
		"src/views/topic/show.html",
	)

	return templates
}
