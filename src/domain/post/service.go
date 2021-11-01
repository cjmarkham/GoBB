package post

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	GetTopicPosts(ctx context.Context, topicID uuid.UUID) ([]Post, error)
}

type service struct {
	repository Repository
}

func ProvideService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
