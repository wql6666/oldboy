package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	filename := "/var/log/yum.log"
	tailConfig := tail.Config{
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
	}
	tails, err := tail.TailFile(filename, tailConfig)
	if err != nil {
		fmt.Println("tail file err", err)
		return
	}
	var msg *tail.Line
	var ok bool
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Println("reopen file's file name", tails.Filename)
			time.Sleep(time.Millisecond*100)
			continue
		}
		fmt.Println(msg)
	}
}
