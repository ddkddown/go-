package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

var wg sync.WaitGroup

func main() {
	consumer, err1 := sarama.NewConsumer([]string{"192.168.11.245:9092"}, nil)
	if err1 != nil {
		fmt.Println("consumer conn err:", err1)
		return
	}
	defer consumer.Close()

	partitions, err2 := consumer.Partitions("testGo")
	if err2 != nil {
		fmt.Println("get partitions err:", err2)
		return
	}

	for _, p := range partitions {
		partitionConsumer, err := consumer.ConsumePartition("testGo", p, sarama.OffsetNewest)
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
