package pkg

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Username struct {
	User     string
	Password string
	Database string
}

func MySQL() (db *sql.DB, e error) {

	username := Username{
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
