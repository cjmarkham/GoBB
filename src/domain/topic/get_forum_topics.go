package topic

import (
	"context"

	"github.com/google/uuid"
)

func (s service) GetForumTopics(ctx context.Context, forumID uuid.UUID) ([]Topic, error) {
	return s.repository.FindByForum(ctx, forumID)
}
