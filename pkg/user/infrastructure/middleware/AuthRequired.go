package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AuthRequired() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Get jwt token from cookie
		cookie := c.Cookies("citywalkerJWT")

		// Check if token is valid
		token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"error": err,
			})
		}

		// Getting claims from token
		claims := token.Claims.(*jwt.StandardClaims)

		// Set in context the userID
		c.Locals("email", claims.Issuer)

		// Pass to handler
		return c.Next()
	}
}
