package repositories

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	Create(gorm.DB)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur UserRepository) Create(db gorm.DB) {

}
