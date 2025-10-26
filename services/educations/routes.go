package educations

import (
	"api/ent"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, client *ent.Client) {
	service := NewEducationController(client)
	router.Get("/educations/:user_id", func(c *fiber.Ctx) error {
		return service.GetUserEducation(c)
	})

	router.Get("/educations/:edu_id", func(c *fiber.Ctx) error {
		return service.GetUserEducation(c)
	})
	router.Post("/educations", func(c *fiber.Ctx) error {
		return service.CreateEducation(c)
	})
}
