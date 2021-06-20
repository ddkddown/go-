package main

import (
	"fmt"
	"os"

	"etcd"
	"kafka"
	"sync"
	"tailf"

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
	wg  sync.WaitGroup
	msg []etcd.CollectEntry
)

func run(cfg *ini.File, wg *sync.WaitGroup, index int) (err error) {
	// 循环读数据
	for {
		line, ok := <-tailf.Tails[index].Lines
		if !ok {
			continue
		}

		if len(line.Text) == 0 {
			continue
		}

		//create message
		kafkaMsg := &sarama.ProducerMessage{}
		kafkaMsg.Topic = msg[index].Topic
		kafkaMsg.Value = sarama.StringEncoder(line.Text)

		kafka.MsgChan <- kafkaMsg
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

	err = etcd.Init([]string{cfg.Section("etcd").Key("address").String()})
	if err != nil {
		fmt.Printf("init etcd failed: %v", err)
		os.Exit(1)
	}

	msg, err := etcd.GetConf(cfg.Section("etcd").Key("collect_key").String())
	if err != nil {
		fmt.Printf("get log_cfg failed: %v", err)
		os.Exit(1)
	}

	//2. 根据配置中的日志路径用tail去收集
	for _, m := range msg {
		err = tailf.Init(m.Path)
		if err != nil {
			fmt.Printf("init tail failed: %v", err)
			os.Exit(1)
		}
	}

	//3. 把日志通过samara发往kafka
	wg.Add(1)
	go kafka.SendMsg(&wg)
	for i := 0; i < len(msg); i++ {
		go run(cfg, &wg, i)
		wg.Add(1)
	}

	wg.Wait()
}
