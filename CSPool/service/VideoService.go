package service

import (
	"errors"
	mysqlmodule "main/dao/mysql"
	redismodule "main/dao/redis"
	"main/model"
	"time"
)

func UploadVideoService(input model.VideoInfo) (vid int64, err error) {
	if mysqlmodule.CheckTitleExist(input.VideoTitle) {
		return 0, errors.New("this title is used")
	}

	vid, err = mysqlmodule.InsertVideo(input)
	if err != nil {
		return 0, err
	}
	if mysqlmodule.GetLevel(input.AuthorName) == 1 {
		mysqlmodule.AuthorizeVideo(vid)
		err = redismodule.InsertVideoTime(vid, input.VideoTime)
		if err != nil {
			return 0, err
		}
		err = redismodule.InsertVideoLike(vid)
		if err != nil {
			return 0, err
		}
	} else {
		err = mysqlmodule.InsertUnderreview(vid)
		if err != nil {
			return 0, err
		}
	}
	return vid, nil
}

func GetVideoByIDService(id int64) (info model.VideoInfo, err error) {
	if !mysqlmodule.CheckVidExist(id) {
		return model.VideoInfo{}, errors.New("invalid video id")
	}
	return mysqlmodule.GetVideoByID(id)
}

func GetVideolistByTimeService() (videolist []model.VideoInfo, err error) {
	idlist, err := redismodule.GetTimeList()
	if err != nil {
		return nil, err
	}

	for _, id := range idlist {
		info, err := GetVideoByIDService(id)
		if err != nil {
			return nil, err
		}
		videolist = append(videolist, info)
	}
	return videolist, nil
}

func GetVideolistByLikeService() (videolist []model.VideoInfo, err error) {
	idlist, err := redismodule.GetLikeList()
	if err != nil {
		return nil, err
	}

	for _, id := range idlist {
		info, err := GetVideoByIDService(id)
		if err != nil {
			return nil, err
		}
		videolist = append(videolist, info)
	}
	return videolist, nil
}

func PublishVideoService(vid int64) (err error) {
	if !mysqlmodule.CheckVidExist(vid) {
		return errors.New("invalid video id")
	}
	if mysqlmodule.CheckVideoStatus(vid) == 1 {
		return errors.New("video is already published")
	}
	mysqlmodule.AuthorizeVideo(vid)
	mysqlmodule.DeleteUnderreview(vid)
	err = redismodule.InsertVideoTime(vid, time.Now())
	if err != nil {
		return err
	}
	err = redismodule.InsertVideoLike(vid)
	if err != nil {
		return err
	}
	return
}

func VoteLikeService(vidstring string, vidint64 int64) (err error) {
	err = mysqlmodule.LikeVideo(vidint64)
	if err != nil {
		return err
	}
	err = redismodule.LikeVideo(vidstring)
	if err != nil {
		return err
	}
	return
}

func VoteDislikeService(vidstring string, vidint64 int64) (err error) {
	err = mysqlmodule.DislikeVideo(vidint64)
	if err != nil {
		return err
	}
	err = redismodule.DislikeVideo(vidstring)
	if err != nil {
		return err
	}
	return
}

func VoteService(info model.VoteInfo) (err error) {
	//video_vote负责记录用户是否点过赞
	//status=0:曾经点过赞或踩,但是取消了(数据库中已经有过记录) 1:点赞 -1:点踩 5:没有记录->要初始化
	var newstatus int
	status, err := mysqlmodule.CheckVoteStatus(info.VoterID, info.VideoIDInt64)
	if err != nil {
		return err
	}
	if status == 5 {
		err = mysqlmodule.VoteInit(info.VoterID, info.VideoIDInt64)
		if err != nil {
			return err
		}
	}
	if info.VoteAction == "like" {
		newstatus = 1
		switch status {
		case 1:
			return errors.New("you have like this video before")
		case 0, 5:
			{
				err = VoteLikeService(info.VideoIDString, info.VideoIDInt64)
			}
		case -1:
			{
				err = VoteLikeService(info.VideoIDString, info.VideoIDInt64)
				err = VoteLikeService(info.VideoIDString, info.VideoIDInt64)
			}
		default:
			return errors.New("wrong status")
		}
	} else if info.VoteAction == "dislike" {
		newstatus = -1
		switch status {
		case 1:
			{
				err = VoteDislikeService(info.VideoIDString, info.VideoIDInt64)
				err = VoteDislikeService(info.VideoIDString, info.VideoIDInt64)
			}
		case 0, 5:
			{
				err = VoteDislikeService(info.VideoIDString, info.VideoIDInt64)
			}
		case -1:
			{
				return errors.New("you have dislike this video before")
			}
		default:
			return errors.New("wrong status")
		}
	} else if info.VoteAction == "cancel" {
		newstatus = 0
		switch status {
		case 1:
			err = VoteDislikeService(info.VideoIDString, info.VideoIDInt64)
		case 0, 5:
			return errors.New("you have not vote this video before")
		case -1:
			err = VoteLikeService(info.VideoIDString, info.VideoIDInt64)
		default:
			return errors.New("wrong status")
		}
	} else {
		return errors.New("wrong action")
	}

	err = mysqlmodule.UpdateVote(info, newstatus)
	return err
}
