package types

import (
	"gorm.io/gorm"
)

// User as represented in the database
type User struct {
	gorm.Model
	Username    string `gorm:"uniqueIndex;type:varchar(32)"`
	DisplayName string `gorm:"type:varchar(100)"`
	Email       string `gorm:"uniqueIndex;type:varchar(100)"`
	Password    string
	Bio         string
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
