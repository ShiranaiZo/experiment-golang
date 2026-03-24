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
	GetAll(*[]models.User) error
	GetUserById(*models.User, string) error
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) Create(user *models.User) error {
	err := r.db.Create(&user).Error

	if err != nil {
		logrus.Errorf("Failed to create user on repository: %v", err)
	}

	return err
}

func (r UserRepository) GetAll(users *[]models.User) error {
	err := r.db.Find(&users).Error

	if err != nil {
		logrus.Errorf("Failed to get all users on repository: %v", err)
	}

	return err
}

func (r UserRepository) GetUserById(user *models.User, userId string) error {
	err := r.db.Where("user_id = ?", []byte(userId)).First(&user).Error

	if err != nil {
		logrus.Errorf("Failed to get user by id '%s' on repository: %v", userId, err)
	}

	return err
}
