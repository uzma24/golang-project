package database

import (
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Instance *sql.DB
var err error

type database struct {
}

func ConnectMySQL(connectionString string) (*sql.DB, error) {
	// Instance, err = sql.Open("mysql", connectionString)
	Instance, err = sql.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/movies?parseTime=true")
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to Database")
	}
	log.Println("Connected to Database!!")
	return Instance, err
}
