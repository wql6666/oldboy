package main

import (
	"bufio"
	"fmt"
	"os"
)

type student struct {
	name  string
	age   int
	score float32
}

func main() {
	test2()
	test()
}

func test() {
	var stu1 = &student{}
	str := "xiaowang 18 89.1"
	fmt.Sscanf(str, "%s %d %f", &stu1.name, &stu1.age, &stu1.score)
	fmt.Println(*stu1)
}

func test2() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read failed err=", err)
		return
	}
	fmt.Println("read succ:", str)

}

