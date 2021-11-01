package post

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	FindByTopic(ctx context.Context, id uuid.UUID) ([]Post, error)
}
