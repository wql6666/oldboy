package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var str = []string{
	"http://www.baidu.com",
	"https://www.taobao.com",
	"http://www.google.com",
}

func main() {
	for _, url := range str {
		c := http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (conn net.Conn, e error) {
					timeout := time.Second
					return net.DialTimeout(network, addr, timeout)
				},
			},
		}

		res, err := c.Head(url)
		if err != nil {
			fmt.Println("head err", err)
		}
		fmt.Println(res.Status)
	}
}
