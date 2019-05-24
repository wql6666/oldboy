package main

import "fmt"

func main() {
	//test2()
	map1 := make(map[string]map[string]string, 100)
	modify(map1)
	for k,v:=range map1{
		fmt.Println(k)
		for k1,v1:=range v{
			fmt.Println(k1,v1)
		}
	}
	fmt.Println(len(map1))
	delete(map1,"xiaoming")
	fmt.Println(map1)

}

func test2() {
	//map1:=make(map[string]map[string]string,100)
	////两成map都需要初始化
	//map1["key1"]=make(map[string]string,100)
	//map1["key1"]["key2"]="hello"
	//fmt.Println(map1)
	//modify(map1)
}

func modify(map1 map[string]map[string]string) {
	_, ok := map1["xiaoming"]
	if !ok {
		map1["xiaoming"] = make(map[string]string)
	}
	map1["xiaoming"]["phone"] = "110"
	fmt.Println(map1)
}
