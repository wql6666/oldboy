package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"

	"time"
)

func main() {

	cli,err:=clientv3.New(clientv3.Config{
		DialTimeout:time.Second,
		Endpoints:[]string{"127.0.0.1:2379"},
	})
	if err!=nil{
		fmt.Println("new etcd client err",err)
		return
	}

	//ctx,_:=context.WithTimeout(context.Background(),time.Second)
	//defer cancel()

	rch:=cli.Watch(context.Background(),"/log/agent")
	for k:=range rch{
		for _,ev:=range k.Events{

			fmt.Println(ev.Kv,ev.Type,ev.PrevKv)
		}
	}

}
