package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mccune1224/hoshii/handlers"
	"github.com/mccune1224/hoshii/store"
	"github.com/mccune1224/hoshii/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbURL = os.Getenv("DATABASE_URL")

func main() {
	app := fiber.New()
	app.Use(logger.New())

	psqlDb, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	psqlDb.AutoMigrate(
		&types.User{},
		&types.Wishlist{},
	)
	if err != nil {
		panic(err)
	}
	handler := handlers.NewHandler(
		psqlDb,
		store.NewPostgreUserStore(psqlDb),
		store.NewPostgreWishlistStore(psqlDb),
	)
	handler.AddRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
