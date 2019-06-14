package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"oldboy/LCS/logagent/config"
	"oldboy/LCS/logagent/tailf"
	"time"
)

var client sarama.SyncProducer

func InitKafka() (err error) {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConfig.Producer.Partitioner = sarama.NewRandomPartitioner
	kafkaConfig.Producer.Return.Successes = true

	kafkaAddrs := make([]string, 1)
	kafkaAddrs = append(kafkaAddrs, config.AppConfig.KafkaAddr)
	fmt.Println(config.AppConfig.KafkaAddr)
	fmt.Println(kafkaAddrs)
	client, err := sarama.NewSyncProducer(kafkaAddrs, kafkaConfig)
	if err != nil {
		fmt.Println("prodecer err", err)
		return
	}
	sendToKafka(client)
	return
}

func sendToKafka(client sarama.SyncProducer) {
	for {
		msg := sarama.ProducerMessage{}
		msg.Topic = tailf.TailMsg.Topic
		msg.Value = sarama.StringEncoder("tailf.TailMsg.Msg")
		fmt.Println("-----",msg.Topic,msg.Value)
		pid, offset, err := client.SendMessage(&msg)
		if err != nil {
			fmt.Println("send msg failed ", err)
			return
		}

		fmt.Println(pid, offset)
		time.Sleep(time.Second)
	}
}
