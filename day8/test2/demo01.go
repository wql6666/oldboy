package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type student struct {
	Name  string
	Age   int
	Score int
}

var file = "/root/桌面/test.txt"

func (stu *student) save() error {

	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("marsh err", err)
		return err
	}
	err = ioutil.WriteFile(file, data, 0777)

	return err
}

func (stu *student) load() error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read file err", err)
		return err
	}
	err = json.Unmarshal(data, stu)

	return err
}
