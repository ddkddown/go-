package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"redis/client"
)

var (
	ctx = context.Background()
)

func main() {
	client.Rdb.ZAdd(ctx, "testSort", &redis.Z{10, "1"}, &redis.Z{9, "2"})
	ret := client.Rdb.ZRange(ctx, "testSort", 0, 10)

	for _, i := range ret.Val() {
		fmt.Println(i)
	}
}
