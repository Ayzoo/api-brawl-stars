package pkg

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/itsbeensolong/api-brawl-stars/utils"
)

func MySQL() (db *sql.DB, e error) {

	username := utils.Username{
		User:     os.Getenv("USER_DATABASE"),
		Password: os.Getenv("PASSWORD_DATABASE"),
		Database: os.Getenv("NAME_DATABASE"),
	}

	conn := fmt.Sprintf("%s:%s@/%s", username.User, username.Password, username.Database)

	db, err := sql.Open("mysql", conn)

	if err != nil {
		panic(err.Error())
	}

	return db, nil
}
