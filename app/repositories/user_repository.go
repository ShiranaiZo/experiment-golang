package repositories

import (
	"github.com/ShiranaiZo/experiment-golang/app/database/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	Create(*models.User) error
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur UserRepository) Create(user *models.User) error {
	err := ur.db.Create(&user).Error

	if err != nil {
		logrus.Errorf("Failed to create user on repository: %v", err)
	}

	return err
}
