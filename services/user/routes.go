package user

import (
	"api/ent"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, client *ent.Client) {
	service := NewUserConteroller(client)
	app.Get("/", func(c *fiber.Ctx) error {
		return service.Home(c)
	})
	app.Post("/user", func(c *fiber.Ctx) error {
		return service.CreateUsers(c)
	})
	app.Put("/user/:id", func(c *fiber.Ctx) error {
		return service.UpdateUser(c)
	})
	app.Post("/email", func(c *fiber.Ctx) error {
		return service.SendEmailNotification(c)
	})
	app.Post("/resume", func(c *fiber.Ctx) error {
		return service.UploadUserResume(c)
	})
}
