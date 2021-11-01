package topic

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	FindByForum(ctx context.Context, id uuid.UUID) ([]Topic, error)
	FindBySlug(ctx context.Context, slug string) (*Topic, error)
}
