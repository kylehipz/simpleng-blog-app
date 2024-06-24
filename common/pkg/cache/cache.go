package cache

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func New() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("DB_PASSWORD")

	addr := fmt.Sprintf("%s:%s", host, port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	return client
}
