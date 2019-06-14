package tailf

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"oldboy/LCS/logagent/config"
	"time"
)

type tailMsg struct {
	Msg   string
	Topic string
}




type TailObj struct {
	Tail  *tail.Tail
	Topic string
}

type TailObjMgr struct {
	TailObjMgrs []*TailObj
}

var tailObjMgr TailObjMgr
var tailObj TailObj

func InitTailf(conf []*config.CollectConf) (err error) {
	if len(conf)==0{
		err=fmt.Errorf("invalid config log collect ,conf%v",conf)
	}
	tailConfig := tail.Config{
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
	}
	for _, v := range conf {

		tailObj.Tail, err = tail.TailFile(v.LogPath, tailConfig)
		if err != nil {
			logs.Error("taif file failed ", v.LogPath, err)
			continue
		}
		tailObj.Topic = v.Topic
		tailObjMgr.TailObjMgrs = append(tailObjMgr.TailObjMgrs, &tailObj)
		go readFromTail(&tailObj, v.ChanSize)

	}

	return

}
var TailMsg tailMsg

func readFromTail(tailObj *TailObj, chanSize int) {

	var msg *tail.Line
	var ok bool
	for {
		msg, ok = <-tailObj.Tail.Lines
		if !ok {

			fmt.Println("reopen file's file name", tailObj.Tail.Filename)
			time.Sleep(time.Millisecond * 100)
			continue

		}
		TailMsg.Msg = msg.Text
		TailMsg.Topic = tailObj.Topic

		//SendMsgToKafka(tailMsg)
	}

}

func SendMsgToKafka(msg tailMsg)  {


}