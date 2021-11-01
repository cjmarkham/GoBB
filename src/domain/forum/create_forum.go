package forum

import "context"

func (s service) CreateForum(ctx context.Context, dto DTO) (*Forum, error) {
	return s.repository.Create(ctx, dto)
}
