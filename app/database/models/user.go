package models

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type User struct {
	UserID    ulid.ULID  `gorm:"primaryKey;not null"`
	Name      string     `gorm:"not null;type:varchar(100)"`
	Address   string     `gorm:"type:text"`
	Email     string     `gorm:"not null;type:varchar(120)"`
	Password  string     `gorm:"not null;type:varchar(120)"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index"`
}
