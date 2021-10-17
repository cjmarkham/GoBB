package home

import "github.com/cjmarkham/GoBB/src/domain/forum"

type Controller struct {
	service forum.Service
}

func ProvideHomeController (service forum.Service) (*Controller, error) {
	return &Controller{
		service: service,
	}, nil
}
