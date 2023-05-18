package db

import (
	"listr/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func RunAutoMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Wishlist{},
	)
}
