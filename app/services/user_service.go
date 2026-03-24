package services

import (
	"github.com/ShiranaiZo/experiment-golang/app/repositories"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	repository repositories.IRepositoryRegistry
}

type IUserService interface {
	Create(*gin.Context)
}

func NewUserService(repository repositories.IRepositoryRegistry) IUserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) Create(ctx *gin.Context) {

}
