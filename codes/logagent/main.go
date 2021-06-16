package main

import (
	"fmt"
	"os"

	"kafka"
	"tailf"

	"sync"

	"github.com/Shopify/sarama"
	"gopkg.in/ini.v1"
)

type KafkaConfig struct {
	Address  string `ini:address`
	Topic    string `ini:topic`
	ChanSize int    `ini:chan_size`
}

type TailConfig struct {
	LogFilePath string `ini:logfile_path`
}

type Config struct {
	KafkaConfig `ini:kafka`
	TailConfig  `ini:collect`
}

var (
	wg sync.WaitGroup
)

func run(cfg *ini.File, wg *sync.WaitGroup) (err error) {
	// 循环读数据
	for {
		line, ok := <-tailf.Tails.Lines
		if !ok {
			continue
		}

		//create message
		msg := &sarama.ProducerMessage{}
		msg.Topic = cfg.Section("kafka").Key("topic").String()
		msg.Value = sarama.StringEncoder(line.Text)

		kafka.MsgChan <- msg
	}

	wg.Done()
	return
}

func main() {
	//0. 读配置文件

	cfg, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Printf("Fail to load cfg: %v", err)
		os.Exit(1)
	}

	/*
		configobj := new(Config)
		err = cfg.MapTo(configobj)
		if err != nil {
			fmt.Printf("Fail to read file: %v", err)
			os.Exit(1)
		}
	*/

	//1. 初始化
	chanSize, _ := cfg.Section("kafka").Key("chan_size").Int()
	err = kafka.Init([]string{cfg.Section("kafka").Key("address").String()}, chanSize)
	if err != nil {
		fmt.Printf("init kafka failed: %v", err)
		os.Exit(1)
	}
	//2. 根据配置中的日志路径用tail去收集
	err = tailf.Init(cfg.Section("collect").Key("logfile_path").String())
	if err != nil {
		fmt.Printf("init tail failed: %v", err)
		os.Exit(1)
	}

	//3. 把日志通过samara发往kafka
	wg.Add(2)
	go kafka.SendMsg(&wg)
	go run(cfg, &wg)

	wg.Wait()
}
