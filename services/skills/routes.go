package skills

import (
	"api/ent"
	"api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, client *ent.Client) {
	service := NewSkillController(client)
	app.Use(middlewares.APIKeyMiddleware())

	// get requests
	app.Get("/skill/:id", func(c *fiber.Ctx) error {
		return service.GetSkill(c)
	})

	app.Get("/skills", func(c *fiber.Ctx) error {
		return service.GetSkills(c)
	})
	app.Get("/techstack", func(c *fiber.Ctx) error {
		return service.GetTechStacks(c)
	})

	// posts
	app.Post("/skill", func(c *fiber.Ctx) error {
		return service.CreateSkill(c)
	})

	app.Post("/techstack", func(c *fiber.Ctx) error {
		return service.CreateTechStack(c)
	})
}
