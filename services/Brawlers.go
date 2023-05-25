package services

import "github.com/gofiber/fiber/v2"

func Brawlers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"example": "YES",
	})
}
