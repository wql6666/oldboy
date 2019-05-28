package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {

	rand.Seed(time.Now().UnixNano())//
}
func main() {
	for i := 1; i <= 10; i++ {
		a := rand.Intn(100) + 1
		fmt.Printf("%d\t", a)
	}
	fmt.Println()
	for i := 1; i <= 10; i++ {
		b := rand.Float32()
		fmt.Printf("%f\t", b)
	}
}
