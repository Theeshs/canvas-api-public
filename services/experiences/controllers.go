package experiences

import (
	"api/ent"
	"api/services/user"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ExperienceController struct {
	handler *ExperienceHandler
}

func NewExperienceController(client *ent.Client) *ExperienceController {
	return &ExperienceController{
		handler: NewExperienceHandler(client),
	}
}

func (ec *ExperienceController) GetUserExperiences(c *fiber.Ctx) error {
	userIdStr := c.Params("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to access user experience",
		})
	}

	_, err = user.NewUserHandler(ec.handler.client).FetchUserByID(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	experiences, err := ec.handler.GenUserExperiences(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No experiences available",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "User experiences retrieved",
		"experiences": experiences,
	})
}

func (ec *ExperienceController) GetUserExperience(c *fiber.Ctx) error {
	expIDStr := c.Params("experience_id")
	expId, err := strconv.Atoi(expIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to access experience",
		})
	}
	experience, err := ec.handler.GenUserExperience(uint(expId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Experience not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "User created successfully",
		"experience": experience,
	})
}

func (ec *ExperienceController) CreateExperience(c *fiber.Ctx) error {
	exp := new(Experience)
	if err := c.BodyParser(exp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	createdExp, err := ec.handler.GenCreateExperience(*exp)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to create experience",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Experience created successfully",
		"experience": createdExp,
	})
}

func (ec *ExperienceController) AddSkillWithExperience(c *fiber.Ctx) error {
	association := new(SkillAssociation)

	if err := c.BodyParser(association); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	if err := ec.handler.GenAddSkillsToExperience(*association); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Skill creating failed",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Skill creating success",
	})

}
