package handler

import (
	"citywalker/pkg/place/application"

	"github.com/gofiber/fiber/v2"
)

func Gets() fiber.Handler {
	return func(c *fiber.Ctx) error {
		city := c.Params("id")
		lng := c.Params("lng")

		// Making database petition
		places, err := application.Gets(&city, &lng)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"places": places,
		})
	}
}
