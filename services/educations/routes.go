package educations

import (
	"api/ent"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, client *ent.Client) {
	app.Get("/educations/:user_id", func(c *fiber.Ctx) error {
		return GetUserEducation(c, client)
	})

	app.Get("/educations/:edu_id", func(c *fiber.Ctx) error {
		return GetUserEducation(c, client)
	})
	app.Post("/educations", func(c *fiber.Ctx) error {
		return CreateEducation(c, client)
	})
}
