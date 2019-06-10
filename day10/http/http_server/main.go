package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", logPanic(Hello))
	http.HandleFunc("/user/login", Login)
	err := http.ListenAndServe("localhost:8888", nil)
	if err != nil {
		fmt.Println("listen err", err)
	}

}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world,welcome")
	panic("panic")

}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello please login")

}

func logPanic(handle http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Printf("%s%v", r.RemoteAddr, err)
			}

		}()
		handle(w, r)
	}

}
