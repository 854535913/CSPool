package mysqlmodule

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // 引入 MySQL 驱动
)

var Sdb *sql.DB

func Init() (err error) {
	// 在 DSN 中添加 parseTime=true 参数
	dsn := "root:Tt112211@tcp(127.0.0.1:3306)/CSPool?parseTime=true"
	Sdb, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return Sdb.Ping()
}
