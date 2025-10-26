package educations

import (
	"api/ent"
	"api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, client *ent.Client) {
	service := NewEducationController(client)
	app.Use(middlewares.APIKeyMiddleware())
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
