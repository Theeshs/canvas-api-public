package projects

import (
	"api/ent"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, client *ent.Client) {
	service := NewProjectConteroller(client)

	// get requests
	app.Get("/projects", func(c *fiber.Ctx) error {
		return service.GetProjects(c)
	})

	// posts
	app.Post("/projects", func(c *fiber.Ctx) error {
		return service.CreateProject(c)
	})
}
