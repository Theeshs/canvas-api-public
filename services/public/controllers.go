package public

import (
	"api/ent"
	"log"

	"github.com/gofiber/fiber/v2"
)

func MyData(c *fiber.Ctx, client *ent.Client) error {

	resp, err := GetPortfolioData(client)
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
