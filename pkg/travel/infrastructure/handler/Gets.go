package handler

import (
	"citywalker/pkg/travel/application"

	"github.com/gofiber/fiber/v2"
)

func Gets() fiber.Handler {
	return func(c *fiber.Ctx) error {

		// User email
		id := c.Locals("email").(string)

		// Making database petition
		travels, err := application.Gets(&id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"travels": travels,
		})
	}
}
