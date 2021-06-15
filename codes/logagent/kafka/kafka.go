package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

var (
	Client sarama.SyncProducer
)

//kafka相关操作
func Init(address []string) (err error) {
	//1. 初始化
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConfig.Producer.Partitioner = sarama.NewRandomPartitioner
	kafkaConfig.Producer.Return.Successes = true

	Client, err = sarama.NewSyncProducer(address, kafkaConfig)

	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}

	return
}
