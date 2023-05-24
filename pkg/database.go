package pkg

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func MySQL() (db *sql.DB, e error) {
	value := os.Getenv("USER_DATABASE")

	db, err := sql.Open("mysql", "root:sebastian123_*@/bralwstars")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(value)
	return db, nil
}
