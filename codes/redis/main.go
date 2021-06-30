package main

import (
	"context"
	"fmt"
	"redis/client"
	"redis/comm"
)

var (
	ctx = context.Background()
)

func main() {

	err := client.Rdb.Set(ctx, "key", "value", 10).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("wang:", val)

	i, _ := comm.GetDistrID()
	fmt.Printf("distriID: %d\n", i)
	i, _ = comm.IncrDistrID()
	fmt.Printf("distriID: %d\n", i)

	/*
		now, _ := client.Lock(ctx, 10)
		time.Sleep(120 * time.Second)
		client.Unlock(ctx, now)
	*/
}
