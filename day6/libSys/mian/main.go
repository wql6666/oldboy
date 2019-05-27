package main

import (
	"fmt"
	"time"
)

type book struct {
	//Id          int
	BookName    string
	Author      string
	CopyNum     int
	PublishDate time.Time
}

type student struct {
	name       string
	certifcate int
	sex        string
	grade      string
}

var booklist []book

func main() {
	var book1 = book{
		"十万个为什么",
		"老舍",
		10,
		time.Now(),
	}
	var button int
	show()
	switch button {
	case 1:
		insertBook(book1)
		show()
	case 2:
		bookQuery()
	case 3:

	case 4:
	case 5:

	}
}

func show()  int{
	button := 0
	fmt.Println("欢迎进入图书管理系统：(请输入相应数字选择功能)")
	fmt.Println("1:书籍入入\n2:书籍查询\n3:学生信息管理\n4:借书\n5:书籍管理")
	fmt.Scan(&button)

	return  button
}


func insertBook(b book) {

	booklist = append(booklist, b)

}

func bookQuery() {
	for _, v := range booklist {
		fmt.Println(v)
	}
}
