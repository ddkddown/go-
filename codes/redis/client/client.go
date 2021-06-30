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
	//这里解锁有个bug，如果a获取了锁，在解锁时，先获取锁的值，发现和自己的值一样，和下一步删除锁的操作不是原子的。
	//要是在这中间锁过期，然后有个b获得了锁，a之后的删锁就可能导致误删。
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

func SafeUnlock(ctx context.Context, timeStamp int64) (ret int) {
	//使用lua脚本实现，因为lua脚本在redis里执行是原子性的，所以就可以解决上诉问题。
	//lua脚本代码
	/*
			if redis.call("get",KEYS[1]) == ARGV[1] then
		    return redis.call("del",KEYS[1])
			else
		    	return 0
			end
	*/

	return 0
}
