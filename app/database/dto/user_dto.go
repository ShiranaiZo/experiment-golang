package dto

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type UserResponse struct {
	UserId ulid.ULID `json:"user_id"`
	UserRequest
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
