package client

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func Lock(ctx context.Context, expireTime int64) (now int64, ret bool) {
	now = time.Now().Unix()
	ret, err := Rdb.SetNX(ctx, "mutex", now, time.Duration(expireTime)*time.Second).Result()

	if !ret {
		fmt.Println("lock mutex failed", err)
		return
	}

	fmt.Println("lock succ")
	return
}

func Unlock(ctx context.Context, timeStamp int64) (ret int) {
	val, _ := Rdb.Get(ctx, "mutex").Result()
	i, _ := strconv.ParseInt(val, 10, 64)
	if timeStamp == i {
		if del, _ := Rdb.Del(ctx, "mutex").Result(); 1 != del {
			fmt.Printf("delete mutex failed: %d", del)
			return -1
		}

		fmt.Println("unlock succ:")
		return 0

	}
	if 0 == i {
		fmt.Println("mutex auto released!")
		return 0
	}

	fmt.Println("timeStamp not eq ", timeStamp, i)

	return 1
}
