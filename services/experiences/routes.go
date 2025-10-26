package experiences

import (
	"api/ent"
	"api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, client *ent.Client) {
	service := NewExperienceController(client)
	app.Use(middlewares.APIKeyMiddleware())
	app.Get("/experience/:user_id", func(c *fiber.Ctx) error {
		return service.GetUserExperiences(c)
	})

	app.Get("/experience/:experience_id", func(c *fiber.Ctx) error {
		return service.GetUserExperience(c)
	})
	app.Post("/experience", func(c *fiber.Ctx) error {
		return service.CreateExperience(c)
	})
	app.Put("/experience/skills", func(c *fiber.Ctx) error {
		return service.AddSkillWithExperience(c)
	})
}
