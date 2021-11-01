package topic

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	GetForumTopics(ctx context.Context, forumID uuid.UUID) ([]Topic, error)
}

type service struct {
	repository Repository
}

func ProvideService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
