package user

type Service interface {
}

type service struct {
	repository Repository
}

func ProvideService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
