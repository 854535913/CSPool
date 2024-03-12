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

func LoginHandler(c *gin.Context) {
	var input model.LoginInfo
	if err := c.ShouldBindJSON(&input); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}
	token, err := service.LoginService(input)
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
