package topic

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	FindByForum(ctx context.Context, forumID uuid.UUID) ([]Topic, error)
}
