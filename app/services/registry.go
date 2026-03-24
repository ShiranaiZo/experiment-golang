package services

import "github.com/ShiranaiZo/experiment-golang/app/repositories"

type Registry struct {
	repository repositories.IRepositoryRegistry
}

type IServiceRegistry interface {
	GetUserService() IUserService
}

func NewServiceRegistry(repository repositories.IRepositoryRegistry) IServiceRegistry {
	return &Registry{
		repository: repository,
	}
}

func (registry *Registry) GetUserService() IUserService {
	return NewUserService(registry.repository)
}
