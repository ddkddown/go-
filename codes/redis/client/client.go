package client

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var Rdb *redis.Client

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
