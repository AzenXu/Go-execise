package dbops

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
	err error
)

func init() {
	db, err = sql.Open("mysql", "root:KFCkfcv*259@tcp(localhost:3306)/video_server?charset=utf8")
	db.Ping()

	if err != nil {
		log.Fatal("数据库未能打开")
	}
}


