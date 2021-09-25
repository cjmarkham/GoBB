package forum

import "context"

type Service interface {
	GetForums(ctx context.Context) ([]Forum, error)
}

type service struct {
	repository Repository
}

func ProvideService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
