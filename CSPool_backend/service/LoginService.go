package service

import (
	mysqlmodule "main/dao/mysql"
	"main/model"
)

func LoginService(inputinfo model.LoginInfo) (token string, err error) {
	exist, err := mysqlmodule.CheckUserExist(mysqlmodule.Sdb, inputinfo.Username)
	if err != nil {
		return "", err
	} else if !exist {
		return "", mysqlmodule.ErrorUserNotExist
	}

	match, err := mysqlmodule.Login(inputinfo)
	if err != nil {
		return "", err
	} else if !match {
		return "", mysqlmodule.ErrorInvalidPassword
	}

	return GenerateTokenService(mysqlmodule.GetID(inputinfo.Username), inputinfo.Username)
}

func CheckLevelService(username string) bool {
	if mysqlmodule.GetLevel(username) > 2 {
		return false
	}
	return true
}
