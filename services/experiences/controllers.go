package experiences

import (
	"api/ent"
	"api/services/user"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUserExperiences(c *fiber.Ctx, client *ent.Client) error {
	userIdStr := c.Params("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to access user experience",
		})
	}

	_, err = user.FetchUserByID(client, uint(userId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	experiences, err := GenUserExperiences(uint(userId), client)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No experiences available",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    experiences,
	})
}

func GetUserExperience(c *fiber.Ctx, client *ent.Client) error {
	expIDStr := c.Params("id")
	expId, err := strconv.Atoi(expIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to access experience",
		})
	}
	experience, err := GenUserExperience(uint(expId), client)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Experience not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    experience,
	})
}

func CreateExperience(c *fiber.Ctx, client *ent.Client) error {
	exp := new(Experience)
	if err := c.BodyParser(exp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	createdExp, err := GenCreateExperience(client, *exp)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to create experience",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Experience created successfully",
		"user":    createdExp,
	})
}

func AddSkillWithExperience(c *fiber.Ctx, client *ent.Client) error {
	association := new(SkillAssociation)

	if err := c.BodyParser(association); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	if err := GenAddSkillsToExperience(*association, client); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Skill creating failed",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Skill creating success",
	})

}
