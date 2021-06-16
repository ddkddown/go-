package kafka

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

var (
	Client  sarama.SyncProducer
	MsgChan chan *sarama.ProducerMessage
)

//kafka相关操作
func Init(address []string, chanSize int) (err error) {
	//1. 初始化
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConfig.Producer.Partitioner = sarama.NewRandomPartitioner
	kafkaConfig.Producer.Return.Successes = true

	Client, err = sarama.NewSyncProducer(address, kafkaConfig)

	if err != nil {
		fmt.Println("producer closed, err:", err)
		return err
	}

	MsgChan = make(chan *sarama.ProducerMessage, chanSize)
	return
}

func SendMsg(wg *sync.WaitGroup) (err error) {
	for {
		select {
		case msg := <-MsgChan:
			pid, offset, err := Client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed, err:", err)
				return errors.New("send msg failed!")
			}

			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		}
	}

	wg.Done()

	return
}
