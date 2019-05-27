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

type test1 interface {
	work()
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
	var a test
	var b test1
	fmt.Printf("%T",a)
	stu1:=student{}
	//stu1:=student{"alan",18,"nan"}
	a=&stu1
	b=&stu1
	b.work()
	a.work()
	a.sleep()
	people1:=people{}
	people1.name="bob"
	a=&people1
	a.work()
	a.sleep()

}
