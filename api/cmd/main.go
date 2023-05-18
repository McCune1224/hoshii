package main

import (
	"listr/db"
	"listr/handler"
	"listr/store"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/gin-gonic/gin"
)

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return ":" + port
}

func main() {
	r := gin.Default()

	psqlDB := db.Connect(os.Getenv("DATABASE_URL"))
	db.RunAutoMigrations(psqlDB)

	apiHandler := handler.NewHandler(
		store.NewUserStore(psqlDB),
		store.NewUserStore(psqlDB),
	)

	apiHandler.RegisterRoutes(r)

	port := GetPort()
	if port != ":3000" {
		gin.SetMode(gin.ReleaseMode)
	}
	r.Run(port)
}
