package route

import (
	"citywalker/pkg/travel/infrastructure/handler"

	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	router.Get("/", handler.Gets())
	router.Post("/", handler.Post())
	router.Post("/expenses", handler.PutExpenses())
	router.Put("/users", handler.PutUsers())
	router.Put("/", handler.Put())
}
