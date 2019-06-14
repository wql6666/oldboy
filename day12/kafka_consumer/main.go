package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	consumer, err := sarama.NewConsumer([]string{"192.168.0.105:9092"}, nil)

	if err != nil {
		fmt.Println("new consumer err", err)
		return
	}
	defer consumer.Close()
	partitions, err := consumer.Partitions("nginx_topic")
	if err != nil {
		fmt.Println("consumer partitions err", err)
	}
	for _, partition := range partitions {
		pc, err := consumer.ConsumePartition("nginx_topic", partition, sarama.OffsetNewest)
		if err != nil {
			fmt.Println("pratition consumer err", err)
			return
		}
		defer pc.AsyncClose()
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {

				fmt.Println(msg.Partition, string(msg.Key), string(msg.Value), msg.Topic, msg.Offset)
			}

		}(pc)

	}
	time.Sleep(time.Hour)
}
