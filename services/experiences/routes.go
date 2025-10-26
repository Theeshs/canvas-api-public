package experiences

import (
	"api/ent"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, client *ent.Client) {
	service := NewExperienceController(client)
	router.Get("/experience/:user_id", func(c *fiber.Ctx) error {
		return service.GetUserExperiences(c)
	})

	router.Get("/experience/:experience_id", func(c *fiber.Ctx) error {
		return service.GetUserExperience(c)
	})
	router.Post("/experience", func(c *fiber.Ctx) error {
		return service.CreateExperience(c)
	})
	router.Put("/experience/skills", func(c *fiber.Ctx) error {
		return service.AddSkillWithExperience(c)
	})
}
