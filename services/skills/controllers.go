package skills

import (
	"api/ent"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateSkill(c *fiber.Ctx, client *ent.Client) error {
	skill := new(Skill)
	if err := c.BodyParser(skill); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}
	createdSkill, err := GenCreateSkill(*skill, client)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to create skill",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Skill created successfully",
		"user":    createdSkill,
	})
}

func GetSkills(c *fiber.Ctx, client *ent.Client) error {
	skills, _ := GenSkills(client)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Skills retrieved successfully",
		"user":    skills,
	})
}

func GetSkill(c *fiber.Ctx, client *ent.Client) error {
	skillIDStr := c.Params("id")
	skillId, err := strconv.Atoi(skillIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to access skills",
		})
	}
	skills, _ := GenSkill(uint(skillId), client)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Skill retrieved successfully",
		"user":    skills,
	})
}
