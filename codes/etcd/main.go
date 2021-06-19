package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	test := 1

	if err != nil {
		fmt.Printf("connect failed: %v, %v", err, test)
		return
	}

	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "show", "time")
	cancel()
	if err != nil {
		fmt.Printf("put msg to etcd failed, err:%v\n", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	rsp, err := cli.Get(ctx, "show")
	cancel()
	if err != nil {
		fmt.Printf("get msg from etcd failed, %v", err)
		return
	}

	for _, ev := range rsp.Kvs {
		fmt.Printf("%s:%s, %v\n", ev.Key, ev.Value, test)
	}

}
