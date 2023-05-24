package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itsbeensolong/api-brawl-stars/pkg/services"
)

func main() {
	app := fiber.New()

	app.Get("/", services.App)

	app.Listen(":3000")
}
