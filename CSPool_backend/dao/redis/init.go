package redismodule

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"main/appconfig"
)

var Rdb *redis.Client

func Init(cfg *appconfig.RedisConfig) (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx := context.Background()

	_, err = Rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return
}

func Close() {
	_ = Rdb.Close()
}
