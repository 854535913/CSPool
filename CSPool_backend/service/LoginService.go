package service

import (
	"errors"
	mysqlmodule "main/dao/mysql"
	"main/model"
)

func LoginService(inputinfo model.LoginInfo) (token string, err error) {
	if !mysqlmodule.CheckUserExist(inputinfo.Username) {
		return "", errors.New("invalid username")
	}

	if !mysqlmodule.Login(inputinfo) {
		return "", errors.New("login failed, invalid username or password")
	}

	return GenerateTokenService(mysqlmodule.GetID(inputinfo.Username), inputinfo.Username)
}

func CheckLevelService(username string) bool {
	if mysqlmodule.GetLevel(username) > 2 {
		return false
	}
	return true
}
