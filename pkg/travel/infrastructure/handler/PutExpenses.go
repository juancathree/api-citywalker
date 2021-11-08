package handler

import (
	expenseDomain "citywalker/pkg/expenses/domain"
	"citywalker/pkg/travel/application"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func PutExpenses() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request map[string]interface{}

		// Convert body to travel
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": err,
			})
		}

		amount, err := strconv.ParseFloat(request["amount"].(string), 64)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": err,
			})
		}

		expense := &expenseDomain.Expense{
			Amount:      amount,
			Description: request["description"].(string),
			IsPersonal:  request["isPersonal"].(bool),
		}

		// Update document in database
		err = application.PutExpenses(c.Locals("email").(string), request["travelID"].(string), expense, request["exists"].(bool))
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
