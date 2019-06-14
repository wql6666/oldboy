package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func main() {
	filename := "/workspace/go/src/oldboy/day11/config/config.txt"
	configer, err := config.NewConfig("ini", filename)
	if err != nil {
		fmt.Println("new config failed err", err)
		return
	}
	serverIP := configer.String("server::ip")
	serverPort, err := configer.Int("server::port")
	if err != nil {
		fmt.Println("parse server port config err=", err)
		return
	}
	clientPort, err := configer.Int("client::port")
	if err != nil {
		fmt.Println("parse client port config err=", err)
		return
	}
	clientIP := configer.String("server::port")
	fmt.Printf("服务器的ip是：%s\t端口是：%d\t客户端的ip是：%s\t端口是：%d\t",
		serverIP, serverPort, clientIP, clientPort)

}
