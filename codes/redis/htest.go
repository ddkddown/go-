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
	client.Rdb.HSet(ctx, "myhash", "key1", "value1", "key2", "value2")

	ret := client.Rdb.HMGet(ctx, "myhash", "key1", "key2")

	for _, i := range ret.Val() {
		fmt.Printf("%s\n", i)
	}
}
