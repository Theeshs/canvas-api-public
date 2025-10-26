package educations

import (
	"api/ent"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, client *ent.Client) {
	service := NewEducationController(client)
	app.Get("/educations/:user_id", func(c *fiber.Ctx) error {
		return service.GetUserEducation(c)
	})

	app.Get("/educations/:edu_id", func(c *fiber.Ctx) error {
		return service.GetUserEducation(c)
	})
	app.Post("/educations", func(c *fiber.Ctx) error {
		return service.CreateEducation(c)
	})
}
