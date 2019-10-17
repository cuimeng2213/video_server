package dbops

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:cm2213@tcp(localhost:3306)/video_server")
	if err != nil {
		fmt.Println("open >>> failed")
		panic(err.Error())
	}
}
