package skills

import (
	"api/ent"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type SkillController struct {
	handler *SkillHandler
}

func NewSkillController(client *ent.Client) *SkillController {
	return &SkillController{
		handler: NewSkillHandler(client),
	}
}

func (sc *SkillController) CreateSkill(c *fiber.Ctx) error {
	skill := new(Skill)
	if err := c.BodyParser(skill); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}
	createdSkill, err := sc.handler.GenCreateSkill(*skill)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to create skill",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Skill created successfully",
		"skill":   createdSkill,
	})
}

func (sc *SkillController) GetSkills(c *fiber.Ctx) error {
	skills, _ := sc.handler.GenSkills()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Skills retrieved successfully",
		"skills":  skills,
	})
}

func (sc *SkillController) GetSkill(c *fiber.Ctx) error {
	skillIDStr := c.Params("id")
	skillId, err := strconv.Atoi(skillIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to access skills",
		})
	}
	skills, _ := sc.handler.GenSkill(uint(skillId))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Skill retrieved successfully",
		"skills":  skills,
	})
}

func (sc *SkillController) GetTechStacks(c *fiber.Ctx) error {
	tech, err := sc.handler.GenTechStack(1)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unable to fetch tech stacks",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Techstack retrieved successfully",
		"techstacks": tech,
	})
}

func (sc *SkillController) CreateTechStack(c *fiber.Ctx) error {
	stackPayload := new(TechStackCreateRequest)
	if err := c.BodyParser(stackPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}
	techstack, err := sc.handler.GenCreateTechStack(*stackPayload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unable to create tech stacks",
			"err":     err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Techstack created successfully",
		"techstacks": techstack,
	})
}
