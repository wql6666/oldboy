package main

import "fmt"

type student struct {
	Name          string `json:"student_name"`
	Age           int    `json:"age"`
	Sex           string `json:"sex"`

}

type people struct {
	student
	//stu1 student
	//stu2 student
	high int
}

func (this *people) work() {
	fmt.Println("people working")
}

func (this *student) init(age int,sex string) {

	this.Age=age
	this.Sex=sex
	//fmt.Println(this)
}

func (this *student) run() {
	fmt.Println("student runing")
}


func main() {
	stu1:=student{"alan",17,"男"}
	stu1.init(100,"女")
	fmt.Println(stu1)
	stu1.run()
	people1:=&people{
		high:100,
	}
	people1.work()
	people1.run()
	people1.Name="alan"
	people1.Age=20
	people1.Sex="nan"
//	people1.stu1.run()
	//people1.stu2.Age=18
	fmt.Println(people1.String())
	fmt.Printf("%s",people1)
}

func (this *people)String()string  {
	str:=fmt.Sprintf("name=%s\tage=%d\tsex=%s\t",this.Name,this.Age,this.Sex)
	return str

}