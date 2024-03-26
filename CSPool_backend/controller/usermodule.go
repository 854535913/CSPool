package controller

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	mysqlmodule "main/dao/mysql"
	"main/model"
	"main/service"
)

func RegisterHandler(c *gin.Context, db *sql.DB) {
	var input model.RegisterInfo
	if err := c.ShouldBindJSON(&input); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}

	if err := service.RegisterService(db, input); err != nil {
		if errors.Is(err, mysqlmodule.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseErrorWithMsg(c, CodeUnexpectedError, err.Error())
		return
	}
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context, db *sql.DB) {
	var input model.LoginInfo
	if err := c.ShouldBindJSON(&input); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}
	token, err := service.LoginService(db, input)
	if err != nil {
		if errors.Is(err, mysqlmodule.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		} else if errors.Is(err, mysqlmodule.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseErrorWithMsg(c, CodeUnexpectedError, err.Error())
		return
	}
	ResponseSuccess(c, token)
}

func GetUseriInfoHandler(c *gin.Context, db *sql.DB) {
	userid, exist := c.Get("UserID")
	if !exist {
		ResponseError(c, CodeCantGetUserID)
		return
	}
	useridInt64, _ := userid.(int64)
	userinfo, err := service.GetUserinfoByID(db, useridInt64)
	if err != nil {
		ResponseErrorWithMsg(c, CodeUnexpectedError, err.Error())
		return
	}
	ResponseSuccess(c, userinfo)
}
