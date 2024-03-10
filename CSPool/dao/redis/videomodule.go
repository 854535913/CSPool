package redismodule

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

func InsertVideoTime(id int64, time time.Time) error {
	ctx := context.Background()
	err := Rdb.ZAdd(ctx, "video_time", &redis.Z{
		Score:  float64(time.Unix()),
		Member: id,
	}).Err()

	if err != nil {
		return err
	}
	return nil
}

func InsertVideoLike(id int64) error {
	ctx := context.Background()
	err := Rdb.ZAdd(ctx, "video_like", &redis.Z{
		Score:  0,
		Member: id,
	}).Err()

	if err != nil {
		return err
	}
	return nil
}

func GetTimeList() (idlist []int64, err error) {
	ctx := context.Background()
	results, err := Rdb.ZRevRangeWithScores(ctx, "video_time", 0, -1).Result()
	if err != nil {
		fmt.Println("Error fetching data from Redis:", err)
		return
	}
	for _, result := range results {
		value, err := strconv.ParseInt(result.Member.(string), 10, 64)
		if err != nil {
			fmt.Println("转换错误：", err)
			return idlist, err
		}
		idlist = append(idlist, value)
	}
	return idlist, nil
}

func GetLikeList() (idlist []int64, err error) {
	ctx := context.Background()
	results, err := Rdb.ZRevRangeWithScores(ctx, "video_like", 0, -1).Result()
	if err != nil {
		fmt.Println("Error fetching data from Redis:", err)
		return
	}
	for _, result := range results {
		value, err := strconv.ParseInt(result.Member.(string), 10, 64)
		if err != nil {
			fmt.Println("转换错误：", err)
			return idlist, err
		}
		idlist = append(idlist, value)
	}
	return idlist, nil
}

func LikeVideo(vid string) (err error) {
	ctx := context.Background()
	err = Rdb.ZIncrBy(ctx, "video_like", 1, vid).Err()
	if err != nil {
		return err
	}
	return
}

func DislikeVideo(vid string) (err error) {
	ctx := context.Background()
	err = Rdb.ZIncrBy(ctx, "video_like", -1, vid).Err()
	if err != nil {
		return err
	}
	return
}
