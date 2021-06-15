package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	//create message
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log")

	client, err := sarama.NewSyncProducer([]string{"192.168.0.102:9092"}, config)

	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()

	//send message
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}

	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
