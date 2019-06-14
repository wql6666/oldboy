package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer([]string{"192.168.0.105:9092"}, config)
	if err != nil {
		fmt.Println("prodecer err", err)
		return
	}
	defer syncProducer.Close()
	for {

	msg := sarama.ProducerMessage{}
	msg.Topic = "nginx_topic"
	msg.Value = sarama.StringEncoder("this is test msg,good!",)
	pid, offset, err := syncProducer.SendMessage(&msg)
	if err != nil {
		fmt.Println("send msg failed ", err)
		return
	}

	fmt.Println(pid, offset, msg.Key, msg.Value)
	time.Sleep(time.Millisecond*100)
	}

}
