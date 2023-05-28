package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itsbeensolong/api-brawl-stars/models"
	"github.com/itsbeensolong/api-brawl-stars/pkg/database"
)

func SearchById(c *fiber.Ctx) error {
	_id := c.Params("id")
	conn, err := database.MySQL()

	if err != nil {
		panic(err.Error())
	}

	query, err := conn.Prepare("SELECT * FROM brawlers WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}
	brawler := models.Brawler{}
	row := query.QueryRow(_id)

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

	err = row.Scan(&id, &name, &image, &types, &category, &stellar, &gadget, &super)

	if err != nil {
		panic(err.Error())
	}

	brawler.Id = id
	brawler.Name = name
	brawler.Image = image
	brawler.Type = types
	brawler.Category = category
	brawler.Stellar = stellars

	return c.JSON(fiber.Map{
		"brawler": brawler,
	})
}
