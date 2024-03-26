package service

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	mysqlmodule "main/dao/mysql"
	redismodule "main/dao/redis"
	"main/model"
	"time"
)

func UploadPostService(sdb *sql.DB, rdb *redis.Client, input model.PostInfo) (pid int64, err error) {
	exist, err := mysqlmodule.CheckTitleExist(sdb, input.PostTitle)
	if err != nil {
		return 0, err
	} else if exist {
		return 0, mysqlmodule.ErrorTitleExist
	}

	pid, err = mysqlmodule.InsertPost(sdb, input)
	if err != nil {
		return 0, err
	}

	level, err := mysqlmodule.GetLevel(sdb, input.AuthorName)
	if err != nil {
		return 0, err
	}
	if level == 1 || level == 2 {
		err = mysqlmodule.AuthorizePost(sdb, pid)

		if err != nil {
			return 0, err
		}
		err = redismodule.InsertPostTime(rdb, pid, input.PostTime)
		if err != nil {
			return 0, err
		}
		err = redismodule.InsertPostLike(rdb, pid)
		if err != nil {
			return 0, err
		}
	} else {
		err = mysqlmodule.InsertUnderreview(sdb, pid)
		if err != nil {
			return 0, err
		}
	}
	return pid, nil
}

func GetPostByIDService(sdb *sql.DB, pid int64) (info model.PostInfo, err error) {
	exist, err := mysqlmodule.CheckPidExist(sdb, pid)
	if err != nil {
		return model.PostInfo{}, err
	} else if !exist {
		return model.PostInfo{}, mysqlmodule.ErrorInvalidPostID
	}
	info, err = mysqlmodule.GetPostByID(sdb, pid)
	return info, err
}

func GetPostlistByTimeService(sdb *sql.DB, rdb *redis.Client) (postlist []model.PostInfo, err error) {
	pidlist, err := redismodule.GetTimeList(rdb)
	if err != nil {
		return nil, err
	}

	for _, id := range pidlist {
		info, err := GetPostByIDService(sdb, id)
		if err != nil {
			return nil, err
		}
		postlist = append(postlist, info)
	}
	return postlist, nil
}

func GetPostlistByLikeService(sdb *sql.DB, rdb *redis.Client) (postlist []model.PostInfo, err error) {
	pidlist, err := redismodule.GetLikeList(rdb)
	if err != nil {
		return nil, err
	}

	for _, id := range pidlist {
		info, err := GetPostByIDService(sdb, id)
		if err != nil {
			return nil, err
		}
		postlist = append(postlist, info)
	}
	return postlist, nil
}

func GetPostlistByUnderreviewService(sdb *sql.DB) (postlist []model.PostInfo, err error) {
	pidlist, err := mysqlmodule.GetPostIDByUnderreview(sdb)
	if err != nil {
		return nil, err
	}

	for _, id := range pidlist {
		info, err := GetPostByIDService(sdb, id)
		if err != nil {
			return nil, err
		}
		postlist = append(postlist, info)
	}
	return postlist, nil
}

func PublishPostService(sdb *sql.DB, rdb *redis.Client, pid int64) (err error) {
	exist, err := mysqlmodule.CheckPidExist(sdb, pid)
	if err != nil {
		return err
	} else if !exist {
		return mysqlmodule.ErrorInvalidPostID
	}

	status, err := mysqlmodule.CheckPostStatus(sdb, pid)
	if err != nil {
		return err
	} else if status == 1 {
		return mysqlmodule.ErrorPostPublished
	}
	err = mysqlmodule.AuthorizePost(sdb, pid)
	if err != nil {
		return err
	}
	err = mysqlmodule.DeleteUnderreview(sdb, pid)
	if err != nil {
		return err
	}
	err = redismodule.InsertPostTime(rdb, pid, time.Now())
	if err != nil {
		return err
	}
	err = redismodule.InsertPostLike(rdb, pid)
	if err != nil {
		return err
	}
	return
}

func VoteLikeService(sdb *sql.DB, rdb *redis.Client, pidstring string, pidint64 int64) (err error) {
	err = mysqlmodule.LikePost(sdb, pidint64)
	if err != nil {
		return err
	}
	err = redismodule.LikePost(rdb, pidstring)
	if err != nil {
		return err
	}
	return
}

func VoteDislikeService(sdb *sql.DB, rdb *redis.Client, pidstring string, pidint64 int64) (err error) {
	err = mysqlmodule.DislikePost(sdb, pidint64)
	if err != nil {
		return err
	}
	err = redismodule.DislikePost(rdb, pidstring)
	if err != nil {
		return err
	}
	return
}

func VoteService(sdb *sql.DB, rdb *redis.Client, info model.VoteInfo) (err error) {
	//video_vote负责记录用户是否点过赞
	//status=0:曾经点过赞或踩,但是取消了(数据库中已经有过记录) 1:点赞 -1:点踩 5:没有记录->要初始化
	var newstatus int
	status, err := mysqlmodule.CheckVoteStatus(sdb, info.VoterID, info.PostIDInt64)
	if err != nil {
		return err
	}
	if status == 5 {
		err = mysqlmodule.VoteInit(sdb, info.VoterID, info.PostIDInt64)
		if err != nil {
			return err
		}
	}
	if info.VoteAction == "like" {
		newstatus = 1
		switch status {
		case 1:
			return mysqlmodule.ErrorVoteStatusLike
		case 0, 5:
			{
				err = VoteLikeService(sdb, rdb, info.PostIDString, info.PostIDInt64)
			}
		case -1:
			{
				err = VoteLikeService(sdb, rdb, info.PostIDString, info.PostIDInt64)
				err = VoteLikeService(sdb, rdb, info.PostIDString, info.PostIDInt64)
			}
		default:
			return mysqlmodule.ErrorVoteStatus
		}
	} else if info.VoteAction == "dislike" {
		newstatus = -1
		switch status {
		case 1:
			{
				err = VoteDislikeService(sdb, rdb, info.PostIDString, info.PostIDInt64)
				err = VoteDislikeService(sdb, rdb, info.PostIDString, info.PostIDInt64)
			}
		case 0, 5:
			{
				err = VoteDislikeService(sdb, rdb, info.PostIDString, info.PostIDInt64)
			}
		case -1:
			{
				return mysqlmodule.ErrorVoteStatusDislike
			}
		default:
			return mysqlmodule.ErrorVoteStatus
		}
	} else if info.VoteAction == "cancel" {
		newstatus = 0
		switch status {
		case 1:
			err = VoteDislikeService(sdb, rdb, info.PostIDString, info.PostIDInt64)
		case 0, 5:
			return mysqlmodule.ErrorVoteStatusNone
		case -1:
			err = VoteLikeService(sdb, rdb, info.PostIDString, info.PostIDInt64)
		default:
			return mysqlmodule.ErrorVoteStatus
		}
	} else {
		return mysqlmodule.ErrorInvalidVoteAction
	}

	err = mysqlmodule.UpdateVote(sdb, info, newstatus)
	return err
}
