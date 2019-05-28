package main

import (
	"fmt"
	"os"
)

//自定义错误
func main() {
	name := "/root/fjak.txt"
	err := openFile(name)
	switch err.(type) {
	case *PathError1:
		fmt.Println("shide", err)
	default:
		fmt.Println("bushide")
	}
}

type PathError1 struct {
	Path    string
	Op      string
	Message error
	Code    int
}

func (p *PathError1) Error() string {
	return fmt.Sprintf("path:%s,op:%s,Message:%v,Code:%d", p.Path, p.Op, p.Message, p.Code)
}

func openFile(name string) error {
	_, err := os.Open(name)
	if err != nil {
		return &PathError1{
			Path:    "haha",
			Op:      "read",
			Message: err,
			Code:    8,
		}
	}
	return nil
}
