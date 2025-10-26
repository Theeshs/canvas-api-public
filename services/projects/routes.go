package projects

import (
	"api/ent"
	"api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, client *ent.Client) {
	service := NewProjectConteroller(client)
	app.Use(middlewares.APIKeyMiddleware())

	// get requests
	app.Get("/projects", func(c *fiber.Ctx) error {
		return service.GetProjects(c)
	})

	// posts
	app.Post("/projects", func(c *fiber.Ctx) error {
		return service.CreateProject(c)
	})
}
