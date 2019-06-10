package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res,err:=http.Get("http://localhost:8888/user/login")
	if err!=nil{
		fmt.Println("get err",err)
	}
	data,err:=ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println("read all err",err)
	}
	fmt.Println(string(data))
}
