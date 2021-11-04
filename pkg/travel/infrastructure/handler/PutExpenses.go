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

		amount := request["expense"].(map[string]interface{})["amount"]
		isPersonal := request["expense"].(map[string]interface{})["isPersonal"]

		amountFloat, err := strconv.ParseFloat(amount.(string), 64)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": err,
			})
		}

		isPersonalBool, err := strconv.ParseBool(isPersonal.(string))

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": err,
			})
		}

		expense := &expenseDomain.Expense{
			Amount:      amountFloat,
			Description: request["expense"].(map[string]interface{})["description"].(string),
			IsPersonal:  isPersonalBool,
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
