package dbops

import "database/sql"

var (
	dbConn *sql.DB
	err error
)

func init(){
	dbConn, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/video?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
