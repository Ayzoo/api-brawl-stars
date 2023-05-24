package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/itsbeensolong/api-brawl-stars/pkg"
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
	db, err := pkg.MySQL()

	if err != nil {
		fmt.Printf("Error obteniendo base de datos: %v", err.Error())
		return
	}

	v1 := app.Group("/api/v1")

	app.Get("/", services.App)

	v1.Get("brawlers", func(c *fiber.Ctx) error {
		return c.SendString("brawlers")
	})
	defer db.Close()

	err = db.Ping()

	if err != nil {
		fmt.Printf("Error conectando: %v", err.Error())
		return
	}
	fmt.Println("Conexion correctamente")
	log.Fatal(app.Listen(*PORT))
}
