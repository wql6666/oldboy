package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type student struct {
	UserId   int    `db:"user_id"`
	UserName string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

var db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:alan0123@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("connect to mysql failed err", err)
		return
	}
	db = database
}

func main() {
	query()
}

func insert() {
	result, err := db.Exec("insert into student1 (username,sex,email) values (?,?,?)", "王二狗", "男", "123@qq.com")
	if err != nil {
		fmt.Println("insert err=", err)
		return
	}
	id, _ := result.LastInsertId()
	fmt.Println("insert succ id=", id)
}

func query() {
	var students []student
	err := db.Select(&students, "select * from student1 where user_id=?", 1)
	if err != nil {
		fmt.Println("select failed err", err)
		return
	}
	fmt.Println(students)
}
