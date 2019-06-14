package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//goroutine 超时控制
type result struct {
	r   *http.Response
	err error
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println("get  new request failed ", err)
		return

	}
	ch1 := make(chan result, 1)
	go func() {
		response, err := client.Do(req)
		if err != nil {
			fmt.Println("request do err", err)
			return
		}
		pack := result{response, err}
		ch1 <- pack
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		result := <-ch1
		fmt.Println(result.err)

	case result := <-ch1:
		defer result.r.Body.Close()
		out, _ := ioutil.ReadAll(result.r.Body)
		fmt.Println(string(out))

	}

}
