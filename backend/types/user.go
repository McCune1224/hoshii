package types

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `json:"name"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Bio         string `json:"bio"`
}

func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		Username:    u.Username,
		DisplayName: u.DisplayName,
		Email:       u.Email,
		Password:    u.Password,
		Bio:         u.Bio,
	}
}

type UserResponse struct {
	Username    string `json:"name"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Bio         string `json:"bio"`
}
