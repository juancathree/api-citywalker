package handler

import (
	"citywalker/pkg/user/application"
	"citywalker/pkg/user/domain"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Login() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var requestBody domain.User

		// Convert body to user
		err := c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": "couldn't be parsed body to user",
			})
		}

		// Check if credentials are right
		err = application.Login(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"error": err,
			})
		}

		// Create jwt token
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    requestBody.Email,
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
		})

		// Sign the token
		token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": "couldn't be signed the jwt token",
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"citywalkerJWT": token,
		})
	}
}
