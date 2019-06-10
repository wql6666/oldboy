package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:alan0123@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("begin action err=", err)
		return
	}
	//defer func() {
	//	recover()
	//}()
	//模拟事务提交
	var aff1, aff2 int64
	result1, err := tx.Exec("update account set money=? where id=?", 2000, 1)
	if result1 != nil {
		fmt.Println("result1 err=", err)
		aff1, _ = result1.RowsAffected()
	}

	result2, err := tx.Exec("update account set money2=? where id=?", 3000, 2)

	if result2 != nil {
		fmt.Println("result2 err=", err)
		aff2, _ = result2.RowsAffected()
	}
	fmt.Println(aff1, aff2)
	if aff1 == 1 && aff2 == 1 {

		tx.Commit()
	} else {
		tx.Rollback()
	}

}
