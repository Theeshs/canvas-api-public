package educations

import (
	"api/ent"
	"api/services/user"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUserEducation(c *fiber.Ctx, client *ent.Client) error {
	userIdStr := c.Params("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to access user educations",
		})
	}

	_, err = user.FetchUserByID(client, uint(userId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	education, err := GenUserEducations(uint(userId), client)
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

func GetUserExperience(c *fiber.Ctx, client *ent.Client) error {
	eduIDStr := c.Params("id")
	eduId, err := strconv.Atoi(eduIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to access experience",
		})
	}
	education, err := GenUserEducation(uint(eduId), client)
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

func CreateEducation(c *fiber.Ctx, client *ent.Client) error {
	edu := new(Education)
	if err := c.BodyParser(edu); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	createdExp, err := GenCreateUserEducation(*edu, client)

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
