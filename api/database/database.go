package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

var Ctx = context.Background()

/*
 * NewClient(): This method takes a struct with the properties to connect to the Redis Server.
 * Addr : This describes the address and port to the Redis server instance.
 * Password : Password to the Redis instance.
 * DB : The database index to use for the application.
*/

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: 		os.Getenv("APP_DB_ADDRESS"),
		Password: 	os.Getenv("APP_DB_PASSWORD"),
		DB: 		dbNo,
	})
	return rdb
}
