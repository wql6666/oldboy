package config

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
)

type AppConf struct {
	LogPath  string
	LogLevel string
	KafkaAddr string
	Collects []*CollectConf
}

type CollectConf struct {
	LogPath   string
	Topic     string
	ChanSize  int
}

var (
	AppConfig     AppConf
	CollectConfig CollectConf
)

func InitConf(confType, fileName string) (err error) {
	configer, err := config.NewConfig(confType, fileName)
	if err != nil {
		fmt.Println("new config failed err", err)
		return
	}
	AppConfig.LogPath = configer.String("appConf::log_path")
	if len(AppConfig.LogPath) == 0 {
		errors.New("appconfig logpath is not exists")
	}
	AppConfig.LogLevel = configer.String("appConf::log_level")
	if len(AppConfig.LogLevel) == 0 {
		AppConfig.LogLevel = "debug"
	}
	CollectConfig.LogPath = configer.String("collect::log_path")
	if len(CollectConfig.LogPath) == 0 {
		errors.New("collect logPath is not exists")
	}
	CollectConfig.Topic = configer.String("collect::topic")
	if len(CollectConfig.Topic) == 0 {
		fmt.Println("topic is not exists")
	}
	fmt.Println(CollectConfig.Topic)
	CollectConfig.ChanSize, err = configer.Int("collect::chan_size")
	if err != nil {
		fmt.Println("chansezi err=", err)
	}
	AppConfig.KafkaAddr = configer.String("kafka::kafka_addr")
	if len(AppConfig.KafkaAddr) == 0 {
		fmt.Println("AppConfig kafkaAddr is not correct")
	}

	AppConfig.Collects = append(AppConfig.Collects, CollectConfig)
	return

}
