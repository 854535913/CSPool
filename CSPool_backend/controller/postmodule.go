package controller

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	mysqlmodule "main/dao/mysql"
	"main/model"
	"main/service"
	"strconv"
	"time"
)

func UploadHandler(c *gin.Context, sdb *sql.DB, rdb *redis.Client) {
	var input model.PostInfo
	if err := c.ShouldBindJSON(&input); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}

	username, exist := c.Get("Username")
	if !exist {
		ResponseError(c, CodeCantGetUsername)
		return
	}
	authorname, _ := username.(string)

	input.AuthorName = authorname
	input.PostTime = time.Now()

	vid, err := service.UploadPostService(sdb, rdb, input)
	if err != nil {
		if errors.Is(err, mysqlmodule.ErrorTitleExist) {
			ResponseError(c, CodeTitleExist)
			return
		}
		ResponseErrorWithMsg(c, CodeUnexpectedError, err.Error())
		return
	}
	ResponseSuccess(c, vid)
}

func GetPostByIDHandler(c *gin.Context, sdb *sql.DB) {
	pidstring := c.Param("id")
	pidint64, _ := strconv.ParseInt(pidstring, 10, 64)
	var info model.PostInfo

	info, err := service.GetPostByIDService(sdb, pidint64)
	if err != nil {
		if errors.Is(err, mysqlmodule.ErrorInvalidPostID) {
			ResponseError(c, CodeInvalidPostID)
			return
		}
		ResponseErrorWithMsg(c, CodeUnexpectedError, err.Error())
		return
	}
	if info.IssueStatus == false {
		ResponseError(c, CodePostNotApprove)
		return
	}
	ResponseSuccess(c, info)
}

func GetPostlistByTimeHandler(c *gin.Context, sdb *sql.DB, rdb *redis.Client) {
	postlist, err := service.GetPostlistByTimeService(sdb, rdb)
	if err != nil {
		ResponseErrorWithMsg(c, CodeUnexpectedError, err.Error())
		return
	}
	ResponseSuccess(c, postlist)
}

func GetPostlistByLikeHandler(c *gin.Context, sdb *sql.DB, rdb *redis.Client) {
	postlist, err := service.GetPostlistByLikeService(sdb, rdb)
	if err != nil {
		ResponseErrorWithMsg(c, CodeUnexpectedError, err.Error())
		return
	}
	ResponseSuccess(c, postlist)
}

func GetPostlistByUnderreviewHandler(c *gin.Context, sdb *sql.DB) {
	postlist, err := service.GetPostlistByUnderreviewService(sdb)
	if err != nil {
		ResponseErrorWithMsg(c, CodeUnexpectedError, err.Error())
		return
	}
	ResponseSuccess(c, postlist)
}

func ReviewHandler(c *gin.Context, sdb *sql.DB, rdb *redis.Client) {
	username, exist := c.Get("Username")
	if !exist {
		ResponseError(c, CodeCantGetUsername)
		return
	}
	pidstring := c.Param("id")
	pidint64, _ := strconv.ParseInt(pidstring, 10, 64)
	usernamestring, _ := username.(string)

	allow, err := service.CheckLevelService(sdb, usernamestring)
	if err != nil {
		ResponseErrorWithMsg(c, CodeUnexpectedError, err.Error())
		return
	} else if allow {
		err = service.PublishPostService(sdb, rdb, pidint64)
		if err != nil {
			if errors.Is(err, mysqlmodule.ErrorInvalidPostID) {
				ResponseError(c, CodeInvalidPostID)
				return
			} else if errors.Is(err, mysqlmodule.ErrorPostPublished) {
				ResponseError(c, CodePostPublished)
				return
			}
			ResponseErrorWithMsg(c, CodeUnexpectedError, err.Error())
			return
		}
		ResponseSuccess(c, "")
	} else {
		ResponseError(c, CodeNeedLevel)
		return
	}
}

func VoteHandler(c *gin.Context, sdb *sql.DB, rdb *redis.Client) {
	userid, exist := c.Get("UserID")
	if !exist {
		ResponseError(c, CodeCantGetUsername)
	}
	useridint64, _ := userid.(int64)
	pidstring := c.Param("id")
	pidint64, _ := strconv.ParseInt(pidstring, 10, 64)
	action := c.Param("action")

	info := model.VoteInfo{
		VoterID:      useridint64,
		PostIDInt64:  pidint64,
		PostIDString: pidstring,
		VoteAction:   action,
		VoteTime:     time.Now(),
	}
	err := service.VoteService(sdb, rdb, info)
	if err != nil {
		if errors.Is(err, mysqlmodule.ErrorVoteStatusLike) {
			ResponseError(c, CodeVoteLike)
			return
		} else if errors.Is(err, mysqlmodule.ErrorVoteStatusDislike) {
			ResponseError(c, CodeVoteDislike)
			return
		} else if errors.Is(err, mysqlmodule.ErrorVoteStatusNone) {
			ResponseError(c, CodeVoteCancel)
			return
		} else if errors.Is(err, mysqlmodule.ErrorVoteStatus) {
			ResponseError(c, CodeVoteInvalidStatus)
			return
		} else if errors.Is(err, mysqlmodule.ErrorInvalidVoteAction) {
			ResponseError(c, CodeVoteInvalidAction)
			return
		}
		ResponseErrorWithMsg(c, CodeUnexpectedError, err.Error())
		return
	}
	ResponseSuccess(c, action)
}
