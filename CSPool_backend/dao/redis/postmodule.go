package redismodule

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

func InsertPostTime(db *redis.Client, id int64, time time.Time) error {
	ctx := context.Background()
	err := db.ZAdd(ctx, "post_time", &redis.Z{
		Score:  float64(time.Unix()),
		Member: id,
	}).Err()

	if err != nil {
		return err
	}
	return nil
}

func InsertPostLike(db *redis.Client, id int64) error {
	ctx := context.Background()
	err := db.ZAdd(ctx, "post_like", &redis.Z{
		Score:  0,
		Member: id,
	}).Err()

	if err != nil {
		return err
	}
	return nil
}

func GetTimeList(db *redis.Client) (idlist []int64, err error) {
	ctx := context.Background()
	results, err := db.ZRevRangeWithScores(ctx, "post_time", 0, -1).Result()
	if err != nil {
		return
	}
	for _, result := range results {
		value, err := strconv.ParseInt(result.Member.(string), 10, 64)
		if err != nil {
			return idlist, err
		}
		idlist = append(idlist, value)
	}
	return idlist, nil
}

func GetLikeList(db *redis.Client) (idlist []int64, err error) {
	ctx := context.Background()
	results, err := db.ZRevRangeWithScores(ctx, "post_like", 0, -1).Result()
	if err != nil {
		return
	}
	for _, result := range results {
		value, err := strconv.ParseInt(result.Member.(string), 10, 64)
		if err != nil {
			return idlist, err
		}
		idlist = append(idlist, value)
	}
	return idlist, nil
}

func LikePost(db *redis.Client, vid string) (err error) {
	ctx := context.Background()
	err = db.ZIncrBy(ctx, "post_like", 1, vid).Err()
	if err != nil {
		return err
	}
	return
}

func DislikePost(db *redis.Client, vid string) (err error) {
	ctx := context.Background()
	err = db.ZIncrBy(ctx, "post_like", -1, vid).Err()
	if err != nil {
		return err
	}
	return
}
