package main

import (
	"fmt"
	"log"
	"uzma/go-backend-clean-architecture/api/route"
	"uzma/go-backend-clean-architecture/config"
	"uzma/go-backend-clean-architecture/database"

	echo "github.com/labstack/echo/v4"
)

func main() {
	conf, err := config.LoadConfig("./config/dev/")
	if err != nil {
		fmt.Println("Error in loading configurations!")
	}

	//connect to MySQL DB
	db, err := database.ConnectMySQL(conf.DatabaseURI)
	if err != nil {
		log.Println("Error connecting to MySQL Database!!")
	}

	//connect to redis
	redisdb, err := database.ConnectRedis(conf.RedisURI)
	if err != nil {
		log.Println("Error connecting to Redis DB!!")
	}

	e := echo.New()
	route.Setup(e, db, redisdb)
	e.Start(":8080") // Start the server on port 8080

}
