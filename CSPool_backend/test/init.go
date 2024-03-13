package test

import (
	"context"
	"database/sql"
	"github.com/go-redis/redis/v8"
)

var sdbTest *sql.DB

func InitMySQL() (err error) {
	// 在 DSN 中添加 parseTime=true 参数
	dsn := "root:Tt112211@tcp(127.0.0.1:3306)/CSPool_test?parseTime=true"
	sdbTest, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return sdbTest.Ping()
}

var rdbTest *redis.Client

func InitRedis() (err error) {
	rdbTest = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis地址
		Password: "",               // 密码，没有则留空
		DB:       1,                // 默认数据库
	})

	ctx := context.Background()
	_, err = rdbTest.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return
}
