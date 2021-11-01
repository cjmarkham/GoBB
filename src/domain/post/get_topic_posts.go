package post

import "context"

func (s service) GetTopicPosts(ctx context.Context, topicID string) ([]Post, error) {
	return s.repository.FindByTopic(ctx, topicID)
}
