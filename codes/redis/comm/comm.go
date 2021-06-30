package comm

import (
	"context"
	"fmt"
	"redis/client"
	"strconv"
	"sync"
)

var (
	ctx  context.Context
	once sync.Once
)

func init() {
	ctx = context.Background()
	client.Rdb.Set(ctx, "distrId", 0, 0)
}

//获取分布式ID
func GetDistrID() (i int, err error) {
	defer func() {
		switch p := recover(); {
		default:
			fmt.Println(p)
		}
	}()

	val, err := client.Rdb.Get(ctx, "distrId").Result()
	if err != nil {
		panic(err)
	}

	i, err = strconv.Atoi(val)

	return
}

//自增分布式ID
func IncrDistrID() (i int, err error) {
	defer func() {
		switch p := recover(); {
		default:
			fmt.Println(p)
		}
	}()

	val, err := client.Rdb.Incr(ctx, "distrId").Result()
	if err != nil {
		panic(err)
	}

	return int(val), err
}
