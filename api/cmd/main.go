package main

import (
	"os"

	"github.com/mccune1224/listr/db"
	"github.com/mccune1224/listr/handler"
	"github.com/mccune1224/listr/store"

	_ "github.com/joho/godotenv/autoload"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return ":" + port
}

func main() {
	app := fiber.New()

	// Add default middlewares
	app.Use(csrf.New())
	app.Use(logger.New())
	app.Use(helmet.New())

	// my lame attempt at preventing csrf attacks
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("FRONTEND_DOMAIN"),
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Connect to db and run migrations
	psqlDB := db.Connect(os.Getenv("DATABASE_URL"))
	db.RunAutoMigrations(psqlDB)

	// Create handler and register routes
	apiHandler := handler.NewHandler(
		store.NewUserStore(psqlDB),
		store.NewUserStore(psqlDB),
	)

	apiHandler.RegisterRoutes(app)

	port := GetPort()
	app.Listen(port)
}
