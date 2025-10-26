package educations

import (
	"api/ent"
	"api/services/user"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type EducationController struct {
	handler *EducationHandler
}

func NewEducationController(client *ent.Client) *EducationController {
	return &EducationController{
		handler: NewEducationHandler(client),
	}
}

func (ec *EducationController) GetUserEducation(c *fiber.Ctx) error {
	userIdStr := c.Params("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to access user educations",
		})
	}

	_, err = user.NewUserHandler(ec.handler.client).FetchUserByID(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	education, err := ec.handler.GenUserEducations(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No educations available",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully rertireved educations",
		"user":    education,
	})
}

func (ec *EducationController) GetUserExperience(c *fiber.Ctx) error {
	eduIDStr := c.Params("id")
	eduId, err := strconv.Atoi(eduIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to access experience",
		})
	}
	education, err := ec.handler.GenUserEducation(uint(eduId), ec.handler.client)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Education not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Education retrieved successfully",
		"user":    education,
	})
}

func (ec *EducationController) CreateEducation(c *fiber.Ctx) error {
	edu := new(Education)
	if err := c.BodyParser(edu); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	createdExp, err := ec.handler.GenCreateUserEducation(*edu, ec.handler.client)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to create education",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Education created successfully",
		"user":    createdExp,
	})
}
