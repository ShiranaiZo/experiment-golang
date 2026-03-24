package controllers

import (
	"github.com/ShiranaiZo/experiment-golang/app/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IControllerRegistry interface {
	GetUserController() IUserController
}

func NewControllerRegistry(service services.IServiceRegistry) IControllerRegistry {
	return &Registry{
		service: service,
	}
}

func (registry *Registry) GetUserController() IUserController {
	return NewUserController(registry.service)
}
