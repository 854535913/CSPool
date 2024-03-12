package test

import "database/sql"

var Sdb_test *sql.DB

func MySQLInit() (err error) {
	// 在 DSN 中添加 parseTime=true 参数
	dsn := "root:Tt112211@tcp(127.0.0.1:3306)/CSPool_test?parseTime=true"
	Sdb_test, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return Sdb_test.Ping()
}
