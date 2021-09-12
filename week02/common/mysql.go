package common

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//创建mysql 连接
func NewMysqlConn() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	return
}
