package service

import (
	"errors"
	mysqlmodule "main/dao/mysql"
	"main/model"
)

func RegisterService(inputinfo model.RegisterInfo) (err error) {
	if mysqlmodule.CheckUserExist(inputinfo.Username) {
		return errors.New("this username is registered")
	}

	mysqlmodule.InsertUser(
		model.UserInfo{
			Username: inputinfo.Username,
			Password: inputinfo.Password,
			Level:    mysqlmodule.InvitationVerify(inputinfo.InvitationCode),
		})
	return
}
