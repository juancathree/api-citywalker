package handler

import (
	"citywalker/pkg/travel/application"
	"citywalker/pkg/travel/domain"

	"github.com/gofiber/fiber/v2"
)

func Post() fiber.Handler {
	return func(c *fiber.Ctx) error {
		lng := c.Params("lng")

		var requestBody domain.Travel

		// Convert body to travel
		err := c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": err,
			})
		}

		// Initialize and save in database
		travel, err := application.Post(&requestBody, c.Locals("email").(string), &lng)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"travel": travel,
		})
	}
}
