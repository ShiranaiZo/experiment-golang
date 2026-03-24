package repositories

import "gorm.io/gorm"

type Registry struct {
	db *gorm.DB
}

type IRepositoryRegistry interface {
	GetUserRepository() IUserRepository
}

func NewRepositoryRegistry(db *gorm.DB) IRepositoryRegistry {
	return &Registry{
		db: db,
	}
}

func (registry *Registry) GetUserRepository() IUserRepository {
	return NewUserRepository(registry.db)
}
