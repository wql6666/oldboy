package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


type emp struct {
	Ename string
	Job string
	Sal float64
}

func main() {
	db, err := sql.Open("mysql", "root:alan0123@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
	//add(db)
	//delete(db)
	//update(db)
	//query(db)
	queryOne(db)
}

func add(db *sql.DB) {
	stmt, err := db.Prepare("insert into emp (empno,ename,job,sal,hiredate)values (?,?,?,?,?)")
	if err != nil {
		fmt.Println("prepare err", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(9535, "王三狗", "读书", 18000, "2019-7-10")
	if err != nil {
		fmt.Println("exec err", err)
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("lastid err=", err)
	}
	rowsAffect, _ := result.RowsAffected()
	fmt.Println("最后的行数=", lastID)
	fmt.Println("影响的行数为", rowsAffect)

}

func delete(db *sql.DB) {

	stmt, err := db.Prepare("delete from emp where empno=?")
	if err != nil {
		fmt.Println("prepare err", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(9534)
	if err != nil {
		fmt.Println("exec err", err)
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("lastid err=", err)
	}
	rowsAffect, _ := result.RowsAffected()
	fmt.Println("最后的行数=", lastID)
	fmt.Println("影响的行数为", rowsAffect)
}
func update(db *sql.DB) {
	result, err := db.Exec("update  emp set job=? where empno=?", "吃饭", 9533)
	if err != nil {
		fmt.Println("exec err", err)
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("lastid err=", err)
	}
	rowsAffect, _ := result.RowsAffected()
	fmt.Println("最后的行数=", lastID)
	fmt.Println("影响的行数为", rowsAffect)
}

func query(db *sql.DB) {
	rows, err := db.Query("select ename,job,sal from emp where deptno=?", 10)
	if err != nil {
		fmt.Println("query err=", err)
		return
	}
	defer rows.Close()
	var (
		ename string
		job   string
		sal   float64
	)
	datas:=make([]emp,0)
	fmt.Println(rows.Columns())
	for rows.Next() {
		rows.Scan(&ename,&job,&sal)
		emp1:=emp{ename,job,sal}
		datas=append(datas,emp1)
	}
	for _,v:=range datas{
		fmt.Println(v)
	}

}

func queryOne(db *sql.DB)  {
	row:=db.QueryRow("select ename,job,sal from emp where deptno=?",10)
	var (
		ename string
		job string
		sal float64
	)
	err:=row.Scan(&ename,&job,&sal)
	if err!=nil{
		fmt.Println("scan err=",err)
		return
	}
	emp1:=emp{ename,job,sal}
	fmt.Println(emp1)

}