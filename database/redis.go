package database

import (
	"log"

	redis "github.com/go-redis/redis/v8"
)

func ConnectRedis(connectionString string) (*redis.Options, error) {
	redisInstance, err := redis.ParseURL(connectionString)
	if err != nil {
		log.Println("Cannot connect to Redis", err)
	}
	log.Println("Connected to Redis!!")
	return redisInstance, nil
}
