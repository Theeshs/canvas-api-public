package user

import (
	"api/ent"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, client *ent.Client) {
	service := NewUserConteroller(client)
	router.Get("/", func(c *fiber.Ctx) error {
		return service.Home(c)
	})
	router.Post("/user", func(c *fiber.Ctx) error {
		return service.CreateUsers(c)
	})
	router.Put("/user/:id", func(c *fiber.Ctx) error {
		return service.UpdateUser(c)
	})
	router.Post("/email", func(c *fiber.Ctx) error {
		return service.SendEmailNotification(c)
	})
	router.Post("/resume", func(c *fiber.Ctx) error {
		return service.UploadUserResume(c)
	})
}
