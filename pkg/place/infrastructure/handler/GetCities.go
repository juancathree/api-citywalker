package handler

import (
	"citywalker/pkg/place/application"

	"github.com/gofiber/fiber/v2"
)

func GetCities() fiber.Handler {
	return func(c *fiber.Ctx) error {

		lng := c.Params("lng")

		// Making database petition
		cities, err := application.GetCities(&lng)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"cities": cities,
		})
	}
}
