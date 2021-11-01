package forum

import "context"

func (s service) GetForums(ctx context.Context) ([]Forum, error) {
	forums, err := s.repository.FindWithParents(ctx)
	if err != nil {
		return nil, err
	}

	return forums, nil
}
