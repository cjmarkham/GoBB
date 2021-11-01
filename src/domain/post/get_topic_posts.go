package post

import (
	"context"

	"github.com/google/uuid"
)

func (s service) GetTopicPosts(ctx context.Context, topicID uuid.UUID) ([]Post, error) {
	return s.repository.FindByTopic(ctx, topicID)
}
