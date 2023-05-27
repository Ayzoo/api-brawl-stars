package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itsbeensolong/api-brawl-stars/pkg/database"
)

func Brawlers(c *fiber.Ctx) error {

	conn, err := database.MySQL()

	if err != nil {
		panic("Hubo un error en la coneccion")
	}

	query, err := conn.Query("SELECT * FROM brawlers")

	if err != nil {
		panic("Hubo un error en la consulta")
	}

	for query.Next() {

	}

	return c.JSON(fiber.Map{
		"Example": "YES",
	})
}
