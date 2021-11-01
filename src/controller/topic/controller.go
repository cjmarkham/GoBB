package topic

import (
	"github.com/cjmarkham/GoBB/src/domain/forum"
	"github.com/cjmarkham/GoBB/src/domain/post"
	"github.com/cjmarkham/GoBB/src/domain/topic"
)

type Controller struct {
	forumService forum.Service
	topicService topic.Service
	postService  post.Service
}

func ProvideController(fs forum.Service, ts topic.Service, ps post.Service) (*Controller, error) {
	return &Controller{
		forumService: fs,
		topicService: ts,
		postService:  ps,
	}, nil
}
