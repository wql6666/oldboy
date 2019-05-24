package main

import "fmt"

type student struct {
	name string
	age int
	score int
	//skill study()
}

func main() {

	var xiaoming *student=&student{
		name:"xiaoming",
		age : 18,
		score:90,
	}
	xiaohong:=new(student)
	var xiaofang student
	xiaofang.name="xiaofang"
	xiaohong.name="xiaohong"
	fmt.Println(xiaoming,xiaohong,xiaofang)
	fmt.Printf("%p\t%p\t%p\t%p",&xiaoming,&xiaoming.name,&xiaoming.age,&xiaoming.score)


}



