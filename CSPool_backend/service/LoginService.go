package service

import (
	"database/sql"
	mysqlmodule "main/dao/mysql"
	"main/model"
)

func LoginService(db *sql.DB, inputinfo model.LoginInfo) (token string, err error) {
	exist, err := mysqlmodule.CheckUserExist(db, inputinfo.Username)
	if err != nil {
		return "", err
	} else if !exist {
		return "", mysqlmodule.ErrorUserNotExist
	}

	match, err := mysqlmodule.Login(db, inputinfo)
	if err != nil {
		return "", err
	} else if !match {
		return "", mysqlmodule.ErrorInvalidPassword
	}

	id, err := mysqlmodule.GetID(db, inputinfo.Username)
	if err != nil {
		return "", err
	}
	return GenerateTokenService(id, inputinfo.Username)
}

func CheckLevelService(db *sql.DB, username string) (allow bool, err error) {
	level, err := mysqlmodule.GetLevel(db, username)
	if err != nil {
		return false, err
	}
	if level > 2 {
		return false, err
	}
	return true, err
}
