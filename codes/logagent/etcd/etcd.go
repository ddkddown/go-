package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type CollectEntry struct {
	Path  string `json:path`
	Topic string `json:topic`
}

var (
	cli *clientv3.Client
)

func Init(address []string) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   address,
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Printf("connect failed: %v", err)
		return err
	}

	return
}

func GetConf(key string) (msg []CollectEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rsp, err := cli.Get(ctx, key)

	if err != nil {
		fmt.Printf("get msg from etcd failed, %v, key:%s", err, key)
		return
	}

	for _, ev := range rsp.Kvs {
		err = json.Unmarshal(ev.Value, msg)
		if err != nil {
			fmt.Printf("Unmarshal cfg msg failed: %s", ev.Value)
			return
		}
	}

	return
}
