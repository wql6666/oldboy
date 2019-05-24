package main

import "fmt"

type student struct {
	name string
	age int
	sex string
}

type people struct {
	student
	height int
}

type test interface {
	work()
	sleep()
}

func (this *student) work() {
	fmt.Println("student work")
}


func (this *student) sleep() {
	fmt.Println("student sleep")
}

func (this *people) sleep() {
	fmt.Println("people sleep")
}

func (this *people) work() {
	fmt.Println("people sleep")
}



func main() {
	var test1 test
	fmt.Printf("%v",test1)
	stu1:=student{"alan",18,"nan"}
	test1=&stu1
	test1.work()
	test1.sleep()
	people1:=people{}
	people1.name="bob"
	test1=&people1
	test1.work()
	test1.sleep()

}
