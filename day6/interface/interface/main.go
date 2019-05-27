package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type book struct {
	Name string
	Id   int
	Num  int
}

type student struct {
	Name string
	age  int
	sex  string
}
type bookArray []book

func (b bookArray) Len() int {
	return len(b)

}

func (b bookArray) Less(i, j int) bool {
	return b[i].Name < b[j].Name
}

func (b bookArray) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func main() {
	var books bookArray
	for i := 1; i <= 10; i++ {
		var b = book{
			fmt.Sprintf("book%d", rand.Intn(100)),
			rand.Intn(100),
			rand.Intn(20),
		}
		books = append(books, b)
	}
	for _,v:=range books {
		fmt.Println(v)
	}
	fmt.Println("-------------------")
	sort.Sort(books)
	for _,v:=range books {
		fmt.Println(v)
	}
}
