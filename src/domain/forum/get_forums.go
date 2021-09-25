package forum

import "context"

func (s service) GetForums(ctx context.Context) ([]Forum, error) {
	return s.repository.FindAll(ctx)
}
