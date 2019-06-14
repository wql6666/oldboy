package log

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"oldboy/LCS/logagent/config"
)

func convertLogLevel(logLevel string) int {
	switch logLevel {
	case "debug":
		return logs.LevelDebug
	case "info":
		return logs.LevelInfo
	case "warn":
		return logs.LevelWarn
	case "error":
		return logs.LevelError

	}
	return logs.LevelDebug

}
func InitLog() (err error) {
	logConf := make(map[string]interface{})
	logConf["filename"] = config.AppConfig.LogPath
	logConf["level"] = convertLogLevel(config.AppConfig.LogLevel)
	logConfig, err := json.Marshal(logConf)
	if err != nil {
		fmt.Println("marshal failed err", err)
		return
	}

	err = logs.SetLogger(logs.AdapterFile, string(logConfig))
	if err != nil {
		fmt.Println("set log err", err)
		return
	}
	return

}
