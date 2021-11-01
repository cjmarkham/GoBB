package forum

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*Forum, error)
	FindBySlug(ctx context.Context, slug string) (*Forum, error)
	FindWithParents(ctx context.Context) ([]Forum, error)
	Create(ctx context.Context, dto DTO) (*Forum, error)
}
