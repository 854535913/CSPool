package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeUnexpectedError
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeCantGetUsername
	CodeServerBusy

	CodeTitleExist
	CodePostNotApprove
	CodeInvalidPostID
	CodePostPublished
	CodeNeedLevel
	CodeVoteLike
	CodeVoteDislike
	CodeVoteCancel
	CodeVoteInvalidStatus
	CodeVoteInvalidAction

	CodeNeedLogin
	CodeTokenFormatError
	CodeInvalidToken
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "成功",
	CodeUnexpectedError: "未知错误",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeCantGetUsername: "无法获取用户名",
	CodeServerBusy:      "服务繁忙",

	CodeTitleExist:        "标题已存在",
	CodePostNotApprove:    "帖子没有获得管理员许可",
	CodeInvalidPostID:     "贴子ID错误",
	CodePostPublished:     "贴子已发布",
	CodeNeedLevel:         "没有审核权限",
	CodeVoteLike:          "已经喜欢这个贴子了",
	CodeVoteDislike:       "已经不喜欢这个贴子了",
	CodeVoteCancel:        "已经取消了",
	CodeVoteInvalidStatus: "状态参数错误",
	CodeVoteInvalidAction: "错误的动作",

	CodeNeedLogin:        "需要登录",
	CodeTokenFormatError: "token格式错误",
	CodeInvalidToken:     "无效的token",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
