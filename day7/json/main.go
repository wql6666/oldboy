package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	testSlice()
}

type user struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func testSlice() {
	var s1 = []map[string]user{}
	map1 := make(map[string]user)
	var user1 = user{
		"alan",
		18,
		"110@qq.com",
	}
	var user2 = user{
		"bob",
		20,
		"119@qq.com",
	}
	map1["man"] = user1
	map1["femal"] = user2
	s1 = append(s1, map1)
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal err=", err)
		return
	}
	fmt.Println(string(data))
	var s2 =[]map[string]user{}
	json.Unmarshal(data,&s2)
	fmt.Println(s2)

}

//func testUnmarshal(data []byte)  {
//
//}