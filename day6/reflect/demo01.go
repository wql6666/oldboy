package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string `json:"student_name"`
	Age  int
	sex  string
}

func (s student) study() {
	fmt.Println(s)
}
func (s student) work() {
	fmt.Println("work")
}

func main() {
	stu1 := student{
		"xiaoming",
		20,
		"nan",
	}
	a := reflect.ValueOf(&stu1).Elem()
	t:=reflect.TypeOf(&stu1).Elem()
	tag:=t.Field(0).Tag.Get("json")
	fmt.Println("tag is :",tag)
	numFiled := a.NumField()
	fmt.Println(numFiled)
	a.Field(0).SetString("lalala")
	for i := 0; i < numFiled; i++ {
		fmt.Println(a.Field(i))
	}

	n:=a.NumMethod()
	fmt.Println("method ",n)
	kind := a.Kind()
	fmt.Println(kind)
	//a.Method(0).Call(nil)

}
