package handler

import (
	"citywalker/pkg/user/application"

	"github.com/gofiber/fiber/v2"
)

func Get() fiber.Handler {
	return func(c *fiber.Ctx) error {

		email := c.Params("email")

		// Making database petition
		user, err := application.Get(&email)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"user": user,
		})
	}
}
