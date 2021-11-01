package post

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]Post, error)
	FindByTopic(ctx context.Context, topicID string) ([]Post, error)
}
