package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/itsbeensolong/api-brawl-stars/services"
)

var (
	PORT = flag.String("port", ":3000", "Port to Listen On")
	PROD = flag.Bool("prod", false, "Enable prefork in production")
)

func main() {
	flag.Parse()
	app := fiber.New(fiber.Config{
		Prefork: *PROD,
	})

	v1 := app.Group("/api/v1")

	app.Get("/", services.App)

	v1.Get("brawlers", func(c *fiber.Ctx) error {
		return c.SendString("brawlers")
	})

	log.Fatal(app.Listen(*PORT))
}
