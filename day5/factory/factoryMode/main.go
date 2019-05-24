package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name          string `json:"student_name"`
	Age           int    `json:"age"`
	Sex           string `json:"sex"`
	StudentLeader `json:"studentLeader"`
	int


}

type StudentLeader struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func main() {
	//var s *student
	//s = newStudent("alan", 18, "男")
	s := &student{}
	s.Name = "alan"
	s.Age = 20
	s.Sex = "nan"
	s.int = 100
	s.StudentLeader.Name = "bob"
	s.StudentLeader.Age = 9
	s.StudentLeader.Score = 90
	fmt.Println(s)
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println("marshal err", err)
		return
	}
	fmt.Println(string(data))
}

func newStudent(name string, age int, sex string) *student {
	return &student{
		Name: name,
		Age:  age,
		Sex:  sex,
	}

}
