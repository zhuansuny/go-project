package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //引入后面的包名而不直接使用这个包中定义的函数，变量等资源
)

func main() {
	db, err := sql.Open("mysql", "root:admin123@/test?charset=utf8")
	checkErr(err)
	defer db.Close()

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("asta", "研发", "2019-9-9")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("id=", id)

	//更新数据

	// db.Prepare函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。
	stmt, err = db.Prepare("update userinfo set username=? where username=?")
	checkErr(err)

	res, err = stmt.Exec("小明", "asta") //stmt.Exec()函数用来执行stmt准备好的SQL语句
	checkErr(err)

	affect, err := res.RowsAffected() //数据库变化的行数
	checkErr(err)

	fmt.Printf("改动的数据有%d个\n", affect)

	//查询数据
	rows, err := db.Query("select * from userinfo") //db.Query()函数用来直接执行Sql返回Rows结果。
	checkErr(err)

	for rows.Next() { //逐次取出数据
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
		fmt.Println()
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println("删除的数据个数", affect)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
