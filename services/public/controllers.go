package public

import (
	"api/ent"
	"log"

	"github.com/gofiber/fiber/v2"
)

type PublicController struct {
	handler *PublicHandler
}

func NewPublicController(client *ent.Client) *PublicController {
	return &PublicController{
		handler: NewPublicHandler(client),
	}
}

func (pc *PublicController) MyData(c *fiber.Ctx) error {

	resp, err := pc.handler.GetPortfolioData()
	if err != nil {
		log.Println("Error retrieving data:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving data",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User retrieved successfully",
		"user":    resp,
	})
}
