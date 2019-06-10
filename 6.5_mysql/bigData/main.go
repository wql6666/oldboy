package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"strconv"
	"sync"
)

func main() {
	db, err := sql.Open("mysql", "root:alan0123@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
	num := 20000
	wg := sync.WaitGroup{}
	var surname = []string{"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈"}
	var lastname = []string{"秋林", "岚岚", "建国", "如花", "二狗"}
	if err != nil {
		fmt.Println("err", err)
	}
	stmt, err := db.Prepare("insert into bigdata values (?,?,?)")
	wg.Add(50)
	for i := 1; i <= 50; i++ {
		go func() {
			for i := 1; i <= num; i++ {

				name := fmt.Sprint(surname[rand.Intn(10)], lastname[rand.Intn(5)])
				stmt.Exec(i, name, strconv.Itoa(rand.Intn(1000000)))

			}
		//	_,err:=db.Exec("commit;")
			if err!=nil{
				fmt.Println("commit err",err)
			}
			wg.Done()
		}()

	}
	wg.Wait()

}
