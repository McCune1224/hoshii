package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	DisplayName string
	Email       string
	Password    string
}
