package redismodule

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func Init() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis地址
		Password: "",               // 密码，没有则留空
		DB:       0,                // 默认数据库
	})

	ctx := context.Background()

	_, err = Rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return
}
