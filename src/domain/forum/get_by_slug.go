package forum

import (
	"context"
)

func (s service) GetBySlug(ctx context.Context, slug string) (*Forum, error) {
	forum, err := s.repository.FindBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	return forum, nil
}
