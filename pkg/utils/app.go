package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Server(port int) {
	app := fiber.New()

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
