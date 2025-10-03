package skills

import (
	"api/ent"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, client *ent.Client) {
	app.Get("/skill/:id", func(c *fiber.Ctx) error {
		return GetSkill(c, client)
	})

	app.Get("/skills", func(c *fiber.Ctx) error {
		return GetSkills(c, client)
	})
	app.Post("/skill", func(c *fiber.Ctx) error {
		return CreateSkill(c, client)
	})
}
