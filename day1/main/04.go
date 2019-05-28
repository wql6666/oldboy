package main

import "fmt"

func main() {
	first, second := 3, 4
	revert(&first, &second)
	fmt.Println(first, second)

}

func revert(m, n *int) {
	*m, *n = *n, *m
	//m,n=n,m(这样改的话就是改变变量的地址，无效)
	return
}
