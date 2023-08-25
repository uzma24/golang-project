package database

import (
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var SqlInstance *sql.DB
var err error

func ConnectMySQL(connectionString string) (*sql.DB, error) {
	SqlInstance, err = sql.Open("mysql", connectionString)
	//SqlInstance, err := sql.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/movies?parseTime=true")
	if err != nil {
		log.Println("Cannot connect to Database", err)
	}
	log.Println("Connected to Database!!")
	return SqlInstance, err
}
