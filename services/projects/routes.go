package projects

import (
	"api/ent"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, client *ent.Client) {
	service := NewProjectConteroller(client)

	// get requests
	router.Get("/projects", func(c *fiber.Ctx) error {
		return service.GetProjects(c)
	})
	router.Get("/projects/:id", func(c *fiber.Ctx) error {
		return service.GetProjectByID(c)
	})

	// posts
	router.Post("/projects", func(c *fiber.Ctx) error {
		return service.CreateProject(c)
	})

	router.Delete("/projects/:id", func(c *fiber.Ctx) error {
		return service.DeleteProjectByID(c)
	})
}
