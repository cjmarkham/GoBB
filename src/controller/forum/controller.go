package forum

import (
	"github.com/cjmarkham/GoBB/src/domain/forum"
	"github.com/cjmarkham/GoBB/src/domain/topic"
)

type Controller struct {
	forumService forum.Service
	topicService topic.Service
}

func ProvideController (fs forum.Service, ts topic.Service) (*Controller, error) {
	return &Controller{
		forumService: fs,
		topicService: ts,
	}, nil
}
