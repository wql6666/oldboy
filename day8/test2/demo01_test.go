package main

import (
	"testing"
	"time"
)

func main() {

}

func TestSave(t *testing.T) {
	var stu1 = student{
		"alan",
		18,
		89,
	}
	err := stu1.save()
	if err != nil {
		t.Fatal("test wrong")
	}

}

func TestLoad(t *testing.T) {
	var stu1 = student{
		"alan",
		18,
		89,
	}
	err := stu1.save()
	if err != nil {
		t.Fatal("test wrong")
	}
	time.Sleep(time.Second*10)
	var stu2 = student{}
	err = stu2.load()
	if err != nil {
		t.Fatal("test load err", err)
	}
	if stu1.Name!=stu2.Name{
		t.Fatal("test load name err")

	}
	if stu2.Age != stu1.Age {
		t.Fatalf("load student failed, age not equal")
	}
	if stu2.Score != stu1.Score {
		t.Fatalf("load student failed, score not equal")
	}

}
