package connection

import (
	_ "github.com/go-sql-driver/mysql"
)

// "root:123456@/mentor_db"

// var SQL_CLIENT *sql.DB

// func init() {
// 	var err error
// 	SQL_CLIENT, err = sql.Open("mysql",
// 		initialize.GLOBAL_CONFIG.Db.Mysql.UserName+":"+initialize.GLOBAL_CONFIG.Db.Mysql.Password+"@tcp("+initialize.GLOBAL_CONFIG.Db.Mysql.Host+")/mentor_db?parseTime=true")
// 	if err != nil {
// 		panic(err)
// 	}
// }
