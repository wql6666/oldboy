package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		env  = os.Getenv("GOOS")
		path = os.Getenv("PATH")
	)

	fmt.Println(env, path)
}
