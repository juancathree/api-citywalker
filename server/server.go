package server

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

var app *fiber.App

func init() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("no env gotten")
	}

	app = fiber.New(fiber.Config{
		Prefork: true,
	})

	// Active console logger
	app.Use(logger.New())

	// Active coors
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Setting up the router
	SetupRouter()

}

func StartServer() {
	_ = app.Listen(os.Getenv("PORT"))
}
