package mysqlmodule

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 引入 MySQL 驱动
	"main/appconfig"
)

var Sdb *sql.DB

func Init(cfg *appconfig.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	Sdb, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return Sdb.Ping()
}

func Close() {
	_ = Sdb.Close()
}
