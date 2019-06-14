package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {
	logMap := make(map[string]interface{})
	logMap["filename"] = "/workspace/go/src/oldboy/day11/log/log.txt"
	logMap["level"] = logs.LevelDebug
	logMarshal, err := json.Marshal(logMap)
	if err != nil {
		fmt.Println("marshal failed err", err)
		return
	}

	err = logs.SetLogger(logs.AdapterFile, string(logMarshal))
	if err != nil {
		fmt.Println("set log err", err)
		return
	}
	logs.Debug("succ init log")
	logs.Info("info log")

}
