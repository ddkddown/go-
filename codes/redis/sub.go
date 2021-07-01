package main

import (
	"context"

	"redis/client"
)

var (
	ctx = context.Background()
)

func main() {

	client.Rdb.Publish(ctx, "test", "msg")
}
