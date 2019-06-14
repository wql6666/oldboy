package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second,
	})
	defer client.Close()
	if err != nil {
		fmt.Println("new etcd client err=", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	_, err = client.Put(ctx, "/log/agent", "8888")
	if err != nil {
		fmt.Println("put err", err)
		return
	}
	cancel()
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	getRes, err := client.Get(ctx, "/log/agent")
	if err != nil {
		fmt.Println("get err", err)
		return
	}
	cancel()
	for _, kv := range getRes.Kvs {
		fmt.Println(string(kv.Key), string(kv.Value))
		//
	}

}
