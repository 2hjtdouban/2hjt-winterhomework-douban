package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	dB, err := sql.Open("mysql","root:qqah3288@/users")
	if err != nil{
		panic(err)
	}
	
	db = dB
}