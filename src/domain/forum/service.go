package forum

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	GetForum(ctx context.Context, id uuid.UUID) (*Forum, error)
	GetBySlug(ctx context.Context, slug string) (*Forum, error)
	GetForums(ctx context.Context) ([]Forum, error)
	CreateForum(ctx context.Context, dto DTO) (*Forum, error)
}

type service struct {
	repository Repository
}

func ProvideService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
