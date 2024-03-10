package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/model"
	"main/service"
	"net/http"
	"strconv"
	"time"
)

func UploadHandler(c *gin.Context) {
	var input model.VideoInfo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	username, exist := c.Get("Username")
	if !exist {
		c.JSON(http.StatusOK, gin.H{
			"msg": "cant get username",
		})
	}
	authorname, ok := username.(string)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": "error in get name",
		})
	}

	input.AuthorName = authorname
	input.VideoTime = time.Now()

	vedioid, err := service.UploadVideoService(input)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":     "Upload video success",
		"videoid": vedioid,
	})
}

func GetVideoByIDHandler(c *gin.Context) {
	vidstring := c.Param("id")
	vidint64, err := strconv.ParseInt(vidstring, 10, 64)
	if err != nil {
		fmt.Println("转换错误:", err)
	}
	var info model.VideoInfo
	info, err = service.GetVideoByIDService(vidint64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if info.IssueStatus == false {
		c.JSON(http.StatusOK, gin.H{
			"msg": "this video did not receive administrator approval",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":         "search video success",
		"ID":          info.VideoID,
		"Like":        info.VideoLike,
		"Time":        info.VideoTime,
		"Title":       info.VideoTitle,
		"Description": info.VideoDescription,
		"Link":        info.VideoLink,
		"Tag":         info.VideoTag,
		"AuthorName":  info.AuthorName,
	})
}

func GetVideolistByTimeHandler(c *gin.Context) {
	videolist, err := service.GetVideolistByTimeService()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "search video success",
		"1":   videolist[0],
		"2":   videolist[1],
		"3":   videolist[2],
	})
}

func GetVideolistByLikeHandler(c *gin.Context) {
	videolist, err := service.GetVideolistByLikeService()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "search video success",
		"1":   videolist[0],
		"2":   videolist[1],
		"3":   videolist[2],
	})
}

func ReviewHandler(c *gin.Context) {
	username, exist := c.Get("Username")
	if !exist {
		c.JSON(http.StatusOK, gin.H{
			"msg": "cant get username",
		})
	}
	vidstring := c.Param("id")
	vidint64, err := strconv.ParseInt(vidstring, 10, 64)
	if err != nil {
		fmt.Println("转换错误:", err)
	}

	username2, ok := username.(string)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": "invalid username",
		})
	}
	if service.CheckLevelService(username2) {
		err = service.PublishVideoService(vidint64)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":           "pass",
			"vid published": vidint64,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "you dont have level to review",
		})
	}
}

func VoteHandler(c *gin.Context) {
	useridTemp, exist := c.Get("UserID")
	if !exist {
		c.JSON(http.StatusOK, gin.H{
			"msg": "cant get userid",
		})
	}
	userid, ok := useridTemp.(int64)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": "error in get name",
		})
	}
	vidstring := c.Param("id")
	vidint64, err := strconv.ParseInt(vidstring, 10, 64)
	if err != nil {
		fmt.Println("转换错误:", err)
	}
	action := c.Param("action")

	info := model.VoteInfo{
		VoterID:       userid,
		VideoIDInt64:  vidint64,
		VideoIDString: vidstring,
		VoteAction:    action,
		VoteTime:      time.Now(),
	}
	err = service.VoteService(info)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":       "success",
		"judgement": action,
	})
}
