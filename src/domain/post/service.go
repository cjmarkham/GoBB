package post

import "context"

type Service interface {
	GetPosts(ctx context.Context) ([]Post, error)
	GetTopicPosts(ctx context.Context, topicID string) ([]Post, error)
}

type service struct {
	repository Repository
}

func ProvideService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
