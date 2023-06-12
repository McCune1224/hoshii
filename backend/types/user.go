package types

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Bio:      u.Bio,
	}
}

type UserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}
