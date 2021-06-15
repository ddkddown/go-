package main

import (
	"fmt"
	"os"

	"kafka"
	"tailf"

	"github.com/Shopify/sarama"
	"gopkg.in/ini.v1"
)

type KafkaConfig struct {
	Address string `ini:address`
	Topic   string `ini:topic`
}

type TailConfig struct {
	LogFilePath string `ini:logfile_path`
}

type Config struct {
	KafkaConfig `ini:kafka`
	TailConfig  `ini:collect`
}

func main() {
	//0. 读配置文件
	configobj := new(Config)
	err := ini.MapTo(configobj, "./conf/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	err = kafka.Init([]string{configobj.Address})
	if err != nil {
		fmt.Printf("init kafka failed: %v", err)
		os.Exit(1)
	}
	//2. 根据配置中的日志路径用tail去收集
	err = tailf.Init(configobj.LogFilePath)
	if err != nil {
		fmt.Printf("init tail failed: %v", err)
		os.Exit(1)
	}

	//3. 把日志通过samara发往kafka
	for {
		line, ok := <-tailf.Tails.Lines
		if !ok {
			continue
		}
		//create message
		msg := &sarama.ProducerMessage{}
		msg.Topic = configobj.Topic
		msg.Value = sarama.StringEncoder(line.Text)

		pid, offset, err := kafka.Client.SendMessage(msg)
		if err != nil {
			fmt.Println("send msg failed, err:", err)
			return
		}

		fmt.Printf("pid:%v offset:%v\n", pid, offset)
	}
}
