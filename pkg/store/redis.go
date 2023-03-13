package store

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// This constant saves the registry stream from Redis.
const CREATE_REGISTRY_STREAM = "create_registry_stream"

func connect() *redis.Client {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file!")
	}

	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_STORE_URL"),
		Password: os.Getenv("REDIS_STORE_PASSWORD"),
	})
}

func Create() {
	rdb := connect()

	val, err := rdb.Do(ctx, "XADD", CREATE_REGISTRY_STREAM, "*", "test", 1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val.(string))
}
