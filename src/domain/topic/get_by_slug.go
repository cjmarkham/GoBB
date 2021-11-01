package topic

import (
	"context"
)

func (s service) GetBySlug(ctx context.Context, slug string) (*Topic, error) {
	return s.repository.FindBySlug(ctx, slug)
}
