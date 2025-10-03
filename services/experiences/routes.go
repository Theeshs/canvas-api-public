package experiences

import (
	"api/ent"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, client *ent.Client) {
	app.Get("/experience/:user_id", func(c *fiber.Ctx) error {
		return GetUserExperiences(c, client)
	})

	app.Get("/experience/:experience_id", func(c *fiber.Ctx) error {
		return GetUserExperience(c, client)
	})
	app.Post("/experience", func(c *fiber.Ctx) error {
		return CreateExperience(c, client)
	})
	app.Put("/experience/skills", func(c *fiber.Ctx) error {
		return AddSkillWithExperience(c, client)
	})
}
