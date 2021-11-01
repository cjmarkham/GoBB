package topic

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	GetForumTopics(ctx context.Context, id uuid.UUID) ([]Topic, error)
	GetBySlug(ctx context.Context, slug string) (*Topic, error)
}

type service struct {
	repository Repository
}

func ProvideService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
