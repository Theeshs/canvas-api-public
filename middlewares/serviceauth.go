package middlewares

import (
	"api/config"

	"github.com/gofiber/fiber/v2"
)

func APIKeyMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		apiKey := c.Get("X-API-KEY")
		expectedKey := config.AppConfig.API_KEY

		if apiKey == "" || apiKey != expectedKey {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or missing API key",
			})
		}

		return c.Next()
	}
}
