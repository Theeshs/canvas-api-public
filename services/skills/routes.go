package skills

import (
	"api/ent"
	"api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, client *ent.Client) {
	service := NewSkillController(client)
	router.Use(middlewares.APIKeyMiddleware())

	// get requests
	router.Get("/skill/:id", func(c *fiber.Ctx) error {
		return service.GetSkill(c)
	})

	router.Get("/skills", func(c *fiber.Ctx) error {
		return service.GetSkills(c)
	})
	router.Get("/techstack", func(c *fiber.Ctx) error {
		return service.GetTechStacks(c)
	})

	// posts
	router.Post("/skill", func(c *fiber.Ctx) error {
		return service.CreateSkill(c)
	})

	router.Post("/techstack", func(c *fiber.Ctx) error {
		return service.CreateTechStack(c)
	})
}
