package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

var Engine = html.New("./pages", ".html")

func Admin(c *fiber.Ctx) error {

	return c.Render("index", fiber.Map{
		"Title":   "Admin",
		"License": "Johan Sebastian",
	})
}
