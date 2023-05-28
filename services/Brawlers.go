package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itsbeensolong/api-brawl-stars/models"
	"github.com/itsbeensolong/api-brawl-stars/pkg/database"
)

func Brawlers(c *fiber.Ctx) error {

	conn, err := database.MySQL()

	if err != nil {
		panic("Hubo un error en la coneccion")
	}

	brawler := models.Brawler{}
	brawlerArray := []models.Brawler{}

	query, err := conn.Query("SELECT * FROM brawlers")

	for query.Next() {
		var id int
		var (
			name     string
			image    string
			types    string
			category string
		)
		var (
			stellar int
			gadget  int
			super   int
		)
		var stellars [2]string

		err = query.Scan(&id, &name, &image, &types, &category, &stellar, &gadget, &super)

		if err != nil {
			panic(err.Error())
		}
		brawler.Id = id
		brawler.Name = name
		brawler.Image = image
		brawler.Type = types
		brawler.Category = category
		brawler.Stellar = stellars

		brawlerArray = append(brawlerArray, brawler)
	}

	if err != nil {
		panic("Hubo un error en la consulta")
	}

	return c.JSON(fiber.Map{
		"brawlers": brawlerArray,
	})
}
