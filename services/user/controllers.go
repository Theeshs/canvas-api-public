package user

import (
	"api/common"
	"api/ent"
	"api/utils"
	"context"
	"fmt"
	"strconv"

	"api/config"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx, client *ent.Client) error {
	return c.SendString("Hello, World!")
}

func CreateUsers(c *fiber.Ctx, client *ent.Client) error {

	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	if CheckUserAvailability(user.Email, client) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User account already available",
		})
	}

	hashedPassword, err := common.GenPasswordHash(user.Password)

	createdUser, err := CreateUser(user.Email, hashedPassword, user.UserName, client)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unknown Error. Please refer to the logs",
		})
	}

	fmt.Printf("Received user: %+v\n", user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    createdUser,
	})

}

func UpdateUser(c *fiber.Ctx, client *ent.Client) error {
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to update user",
		})
	}

	_, err = client.User.Get(context.Background(), uint(idInt))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	updatedUser, err := UpdateUsers(uint(idInt), *user, client)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User update fialed",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    updatedUser,
	})

}

func SendEmailNotification(c *fiber.Ctx, client *ent.Client) error {
	emailMessage := new(EmailMessage)

	if err := c.BodyParser(emailMessage); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email sending failed",
		})
	}
	go utils.SendContactEmail(
		config.AppConfig.OwnerEmail,
		config.AppConfig.EmailPassword,
		config.AppConfig.OwnerEmail,
		emailMessage.Name,
		emailMessage.UserEmail,
		emailMessage.Message,
	)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Sucessfully sent the email",
	})
}
