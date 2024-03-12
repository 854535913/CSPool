package service

import (
	"database/sql"
	mysqlmodule "main/dao/mysql"
	"main/model"
)

func RegisterService(db *sql.DB, inputinfo model.RegisterInfo) (err error) {
	exist, err := mysqlmodule.CheckUserExist(db, inputinfo.Username)
	if err != nil {
		return err
	} else if exist {
		return mysqlmodule.ErrorUserExist
	}
	level, err := mysqlmodule.InvitationVerify(db, inputinfo.InvitationCode)
	if err != nil {
		return err
	}
	err = mysqlmodule.InsertUser(db,
		model.UserInfo{
			Username: inputinfo.Username,
			Password: inputinfo.Password,
			Level:    level,
		})
	return err
}
