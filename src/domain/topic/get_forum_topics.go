package topic

import (
	"context"

	"github.com/google/uuid"
)

func (s service) GetForumTopics(ctx context.Context, id uuid.UUID) ([]Topic, error) {
	return s.repository.FindByForum(ctx, id)
}
