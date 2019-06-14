package main

import (
	"fmt"
	"gopkg.in/olivere/elastic.v2"
)

type student struct {
	Name string
	Age  int
}

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://192.168.0.105:9200"))
	if err != nil {
		fmt.Println("new elasticsearch client err", err)
		return
	}
	stu1 := student{
		"alan",
		18,
	}

	for i := 1; i <= 1000; i++ {
		_, err = client.Index().Index("student").Type("stu").Id(fmt.Sprintf("id:%d",i)).BodyJson(stu1).Do()
		if err != nil {
			fmt.Println("insert elastic err", err)
			return
		}
	}


	fmt.Println("insert succ")

}
