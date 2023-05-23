package db

import (
	"github.com/mccune1224/listr/models"

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

func RunAutoMigrations(db *gorm.DB, clean ...bool) {
	// db.Migrator().DropTable(
	// 	&models.User{},
	// 	&models.Follow{},
	// 	&models.Wishlist{},
	// )
	db.AutoMigrate(
		&models.User{},
		&models.Follow{},
		&models.Wishlist{},
	)
}
