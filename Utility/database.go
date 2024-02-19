package Utility

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//https://github.com/go-sql-driver/mysql MySql驅動
//https://github.com/mattn/go-sqlite3 驅動
//https://github.com/mikespook/mymysql 驅動

func StartConnDatabase() {
	db, err := sql.Open("mysql", "root:1qaz@WSX@/test?charset=utf8")
	checkErr(err)

	//插入資料
	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("astaxie", "研發部門", "2012-12-09")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	//查詢資料
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//刪除資料
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	res, err = stmt.Exec(id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/*
sql.Register：這個存在於 database/sql 的函式是用來註冊資料庫驅動
driver.Driver：Driver 是一個數據函式庫驅動的介面，這個Driver 只能應用在一個 goroutine 裡面
driver.Conn：Conn 是一個數據函式庫連線的介面定義，他定義了一系列方法，這個 Conn 只能應用在一個 goroutine 裡面
driver.Stmtl：Stmt 是一種準備好的狀態，和 Conn 相關聯，而且只能應用於一個 goroutine 中。
driver.Tx：交易處理一般就兩個過程，提交或者回復 (Rollback)。
driver.Execer：這是一個 Conn 可選擇實現的介面，如果這個介面沒有定義，那麼在呼叫 DB.Exec，就會首先呼叫 Prepare 回傳 Stmt，然後執行 Stmt 的 Exec，然後關閉 Stmt。
driver.Result：這個是執行 Update/Insert 等操作回傳的結果介面定義
一 LastInsertId 函式回傳由資料庫執行插入操作得到的自增 ID 號。
一 RowsAffected 函式回傳 query 操作影響的資料條目數。
driver.Rows：Rows 是執行查詢回傳的結果集介面定義
driver.Value：Value 其實就是一個空介面，他可以容納任何的資料
driver.ValueConverter：ValueConverter 介面定義了如何把一個普通的值轉化成 driver.Value 的介面
driver.Valuer：Valuer 介面定義了回傳一個 driver.Value 的方式
database/sql：database/sql 在 database/sql/driver 提供的介面基礎上定義了一些更高階的方法，用以簡化資料庫操作
*/

/*
DSN:Data Source Name支援格式
user@unix(/path/to/socket)/dbname?charset=utf8
user:password@tcp(localhost:5555)/dbname?charset=utf8
user:password@/dbname
user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
*/
