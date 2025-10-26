package public

import (
	"api/ent"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, client *ent.Client) {
	service := NewPublicController(client)
	app.Get("/portfolio", func(c *fiber.Ctx) error {
		return service.MyData(c)
	})
}
