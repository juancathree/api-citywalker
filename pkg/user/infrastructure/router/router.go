package router

import (
	"citywalker/pkg/user/infrastructure/handler"

	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	router.Post("/login", handler.Login())
	router.Post("/signup", handler.Signup())
	router.Get("/user/:email", handler.Get())
}
