package forum

import (
	"context"

	"github.com/google/uuid"
)

func (s service) GetForum(ctx context.Context, forumID uuid.UUID) (*Forum, error) {
	forum, err := s.repository.FindByID(ctx, forumID)
	if err != nil {
		return nil, err
	}

	return forum, nil
}
