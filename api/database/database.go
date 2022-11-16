package database

import(
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("APP_DB_ADDRESS"),
		Password: os.Getenv("APP_DB_PASSWORD"),
		DB: dbNo,
	})
	return rdb
}
