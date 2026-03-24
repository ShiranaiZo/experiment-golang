package services

import (
	"fmt"

	"github.com/ShiranaiZo/experiment-golang/app/database/dto"
	"github.com/ShiranaiZo/experiment-golang/app/database/models"
	"github.com/ShiranaiZo/experiment-golang/app/repositories"
	"github.com/oklog/ulid/v2"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	repository repositories.IRepositoryRegistry
}

type IUserService interface {
	Create(*dto.UserRequest) (dto.UserResponse, error)
	Index() ([]dto.UserResponse, error)
	Show(string) (dto.UserResponse, error)
}

func NewUserService(repository repositories.IRepositoryRegistry) IUserService {
	return &UserService{
		repository: repository,
	}
}

func (s UserService) Create(data *dto.UserRequest) (dto.UserResponse, error) {
	user := models.User{
		UserID:   ulid.Make().String(),
		Name:     data.Name,
		Address:  data.Address,
		Email:    data.Email,
		Password: data.Password,
	}

	err := s.repository.GetUserRepository().Create(&user)

	if err != nil {
		logrus.Errorf("Failed to create user on service: %v", err)
	}

	return dto.UserResponse{
		UserId:      user.UserID,
		UserRequest: *data,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		DeletedAt:   user.DeletedAt,
	}, err
}

func (s UserService) Index() ([]dto.UserResponse, error) {
	users := []models.User{}
	err := s.repository.GetUserRepository().GetAll(&users)

	var data []dto.UserResponse

	for _, user := range users {
		data = append(data, dto.UserResponse{
			UserId: user.UserID,
			UserRequest: dto.UserRequest{
				Name:     user.Name,
				Address:  user.Address,
				Email:    user.Email,
				Password: user.Password,
			},
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		})
	}

	return data, err
}

func (s UserService) Show(userId string) (dto.UserResponse, error) {
	user := models.User{}

	err := s.repository.GetUserRepository().GetUserById(&user, userId)

	fmt.Println(user.Name, "halo")

	return dto.UserResponse{
		UserId: user.UserID,
		UserRequest: dto.UserRequest{
			Name:    user.Name,
			Address: user.Address,
			Email:   user.Email,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}, err
}
