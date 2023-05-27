package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/itsbeensolong/api-brawl-stars/services"
	"github.com/joho/godotenv"
)

var (
	PORT = flag.String("port", ":8000", "Port to Listen On")
	PROD = flag.Bool("prod", false, "Enable prefork in production")
)

func main() {

	flag.Parse()
	godotenv.Load()

	app := fiber.New(fiber.Config{
		Prefork: *PROD,
		Views:   services.Engine,
	})

	v1 := app.Group("/api/v1")

	app.Get("/", services.App)

	v1.Get("/brawlers", services.Brawlers)
	v1.Get("/admin", services.Admin)

	log.Fatal(app.Listen(*PORT))
}
