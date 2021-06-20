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

	if err != nil {
		fmt.Printf("connect failed: %v", err)
		return
	}

	defer cli.Close()

	rch := cli.Watch(context.Background(), "show")

	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type:%s, Key:%s, Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
