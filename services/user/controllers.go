package user

import (
	"api/common"
	"api/ent"
	"api/ent/document"
	"api/utils"
	"context"
	"fmt"
	"strconv"

	"api/config"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	handler *UserHandler
}

func NewUserConteroller(client *ent.Client) *UserController {
	return &UserController{
		handler: NewUserHandler(client),
	}
}

func (uc *UserController) Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func (uc *UserController) CreateUsers(c *fiber.Ctx) error {

	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	if uc.handler.CheckUserAvailability(user.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User account already available",
		})
	}

	hashedPassword, err := common.GenPasswordHash(user.Password)

	createdUser, err := uc.handler.CreateUser(user.Email, hashedPassword, user.UserName)

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

func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
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

	_, err = uc.handler.client.User.Get(context.Background(), uint(idInt))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	updatedUser, err := uc.handler.UpdateUsers(uint(idInt), *user)
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

func (uc *UserController) SendEmailNotification(c *fiber.Ctx) error {
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

func (uc *UserController) UploadUserResume(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("File not found")
	}

	// Open file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Upload to Google Drive
	fileID, err := uc.handler.GenUploadToDrive(file.Filename, src, 1, document.DocumentTypeResume)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "File uploaded to Google Drive successfully",
		"file_id": fileID,
	})
}
