package home

import "github.com/cjmarkham/GoBB/src/domain/forum"

type Controller struct {
	service forum.Service
}

func ProvideController (service forum.Service) (*Controller, error) {
	return &Controller{
		service: service,
	}, nil
}
