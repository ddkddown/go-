package main

import (
	"context"
	"fmt"

	"redis/client"
)

var (
	ctx = context.Background()
)

func main() {
	client.Rdb.LPush(ctx, "testIndex", 1, 2, 3, 4)

	ret := client.Rdb.LRange(ctx, "testIndex", 0, 10)

	for _, i := range ret.Val() {
		fmt.Println(i)
	}
}
