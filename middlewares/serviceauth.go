package middlewares

import (
	"api/config"

	"github.com/gofiber/fiber/v2"
)

// APIKeyMiddleware checks for a valid API key in the request headers
func APIKeyMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		apiKey := c.Get("X-API-KEY")            // get API key from header
		expectedKey := config.AppConfig.API_KEY // the key stored in env

		if apiKey == "" || apiKey != expectedKey {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or missing API key",
			})
		}

		return c.Next() // proceed if valid
	}
}
