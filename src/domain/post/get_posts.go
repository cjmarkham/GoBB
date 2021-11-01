package post

import "context"

func (s service) GetPosts(ctx context.Context) ([]Post, error) {
	return s.repository.FindAll(ctx)
}
