package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"oldboy/day7/balance/balancer"
	"os"
	"time"
)

func main() {
	var Insts []*balancer.Instance
	for i := 0; i < 5; i++ {

		inst := &balancer.Instance{
			Host: fmt.Sprintf("192.168.%d.%d", rand.Intn(256), rand.Intn(256)),
			Port: 8080,
		}
		Insts = append(Insts, inst)
	}
	var input balancer.Balancer
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	_, input = balancer.UseMgrBalance(string(line))
	if input == nil {
		return
	}
	//if string(line) == "random" {fmt.Println(balancer.Mgr.AllBalancer)
	//	fmt.Println("random 模式")
	//	input = &balancer.Random{}
	//} else if string(line) == "round" {
	//	fmt.Println("round 模式")
	//
	//	input = &balancer.RoundRobin{}
	//}
	fmt.Println(balancer.Mgr.AllBalancer)

	for {

		err, inst := input.DoBalance(Insts)
		if err != nil {
			fmt.Println("do balance err", err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}
