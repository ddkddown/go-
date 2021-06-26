package main

import (
	"context"
	"fmt"
	"time"

	"redis/client"
)

var (
	ctx = context.Background()
)

func main() {

	err := client.Rdb.Set(ctx, "key", "value", 10).Err()
	if err != nil {
		panic(err)
	}

	time.Sleep(12 * time.Second)

	val, err := client.Rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("wang:", val)
}
