package handler

import (
	"citywalker/pkg/travel/application"

	"github.com/gofiber/fiber/v2"
)

func PutUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request map[string]interface{}

		// Convert body to travel
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": err,
			})
		}

		// Update document in database
		err = application.PutUsers(request["email"].(string), request["travelID"].(string), request["exists"].(bool))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": "ok",
		})
	}
}
