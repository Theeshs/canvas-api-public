package user

import (
	"api/ent"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, client *ent.Client) {
	app.Get("/", func(c *fiber.Ctx) error {
		return Home(c, client)
	})
	app.Post("/user", func(c *fiber.Ctx) error {
		return CreateUsers(c, client)
	})
	app.Put("/user/:id", func(c *fiber.Ctx) error {
		return UpdateUser(c, client)
	})
	app.Post("/email", func(c *fiber.Ctx) error {
		return SendEmailNotification(c, client)
	})
}
