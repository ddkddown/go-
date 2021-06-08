package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err1 := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)

	if err1 != nil {
		fmt.Println("create producer failed", err1)
		return
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic:     "testGo",
		Partition: int32(-1),
		Key:       sarama.StringEncoder("key"),
	}

	msg.Value = sarama.ByteEncoder("message from go client")
	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		fmt.Println("send Message Fail", err)
	}

	fmt.Println("Partion = %d, offset = %d\n", partition, offset)
}
