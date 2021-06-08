package main

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

var wg sync.WaitGroup

func main() {
	consumer, err := sarama.NewConsumer([]string{"192.168.0.102:9092"}, nil)
	if err != nil {
		fmt.Println("consumer conn err:", err)
		return
	}
	defer consumer.Close()

	partitions, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Println("get partitions err:", err)
		return
	}

	for _, p := range partitions {
		partitionConsumer, err := consumer.ConsumePartition("web_log", p, sarama.OffsetNewest)
		if err != nil {
			fmt.Println("partitionConsumer err:", err)
			continue
		}

		wg.Add(1)
		go func() {
			for m := range partitionConsumer.Messages() {
				fmt.Printf("key:%v, text:%v, offset:%d\n", m.Key, m.Value, m.Offset)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
