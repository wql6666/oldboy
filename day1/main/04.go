package main

import "fmt"

func main() {
	revert(3,4)
}

func revert(m,n int)  {
	m,n=n,m
	fmt.Println(m,n)

}
