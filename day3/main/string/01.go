package main

import (
	"fmt"
	"strings"
)

func main() {
	url := "www.baidu.com"
	fmt.Println(hasprefix(&url))
	fmt.Println(hassuffix(&url))
	location := strings.Index(url, "www")
	location1 := strings.LastIndex(url, "s")
	fmt.Println(location1)
	fmt.Println(location)
	fmt.Println(strings.Repeat("hah ", 3))
	fmt.Println(strings.Replace("wojwintiwanhawokaiwxiwn", "w", "我", 5))
	fmt.Println(strings.Count("lalalwowowo", "wo"))
	fmt.Println(strings.ToUpper("nike"))
	fmt.Println(strings.TrimSpace("              lal    hahah lll   "))
	fmt.Println("---------")
	fmt.Println(strings.Trim("www.baiwwwdu.com.www", "www"))
	fmt.Println(strings.TrimLeft("www.baiwwwdu.com.www", "www"))
	fmt.Println(strings.TrimRight("www.baiwwwdu.com.www", "www"))
	fmt.Println(strings.Fields("              lal    hahah lll   "))
	fmt.Println(strings.Split("dfafalalhahahllldgaflff", "l"))
	var s1 []string = []string{"a", "b", "c", "d", "e"}
	fmt.Println(strings.Join(s1, "我"))
	str:=fmt.Sprint(1,2,3)
	fmt.Println(str)
	fmt.Printf("%T",str)

}

func hasprefix(url *string) string {

	if !strings.HasPrefix(*url, "http://") {

		*url = "http://" + *url
	}
	return *url

}

func hassuffix(url *string) string {
	if !strings.HasSuffix(*url, "/") {
		*url = *url + "/"
	}
	return *url

}
