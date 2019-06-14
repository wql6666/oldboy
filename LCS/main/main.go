package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"oldboy/LCS/logagent/config"
	"oldboy/LCS/logagent/kafka"
	"oldboy/LCS/logagent/log"
	"oldboy/LCS/logagent/tailf"
)

func main() {
	confType := "ini"
	confFileName := "/workspace/go/src/oldboy/LCS/conf/appConf.txt"
	err := config.InitConf(confType, confFileName)
	if err != nil {
		fmt.Println("init config failed err", err)
	}

	err = log.InitLog()
	if err != nil {
		fmt.Println("init log failed err", err)
	}
	logs.Debug("init conf&log succ")

	err = tailf.InitTailf(config.AppConfig.Collects)

	err = kafka.InitKafka()
	logs.Debug("init all succ")

}
