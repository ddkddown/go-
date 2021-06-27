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
	client.Rdb.SAdd(ctx, "testset", 11)
	ret := client.Rdb.SMembers(ctx, "testset")

	for _, i := range ret.Val() {
		fmt.Println(i)
	}
}
