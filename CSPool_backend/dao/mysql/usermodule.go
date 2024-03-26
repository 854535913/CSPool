package mysqlmodule

import (
	"database/sql"
	"main/model"
)

func CheckUserExist(db *sql.DB, username string) (exist bool, err error) {
	sqlStr := "SELECT EXISTS(SELECT 1 FROM user WHERE username = ?)"
	err = db.QueryRow(sqlStr, username).Scan(&exist)
	return exist, err
}

func InvitationVerify(db *sql.DB, invitationcode string) (level int8, err error) {
	var exists bool
	sqlStr := "SELECT EXISTS(SELECT 1 FROM invitation WHERE code = ?)"
	err = db.QueryRow(sqlStr, invitationcode).Scan(&exists)
	if err != nil {
		return 0, err
	}
	if exists {
		return 2, nil
	}
	return 3, nil
}

func InsertUser(db *sql.DB, info model.UserInfo) (err error) {
	sqlStr := "insert into user(username, password, level) values(?,?,?)"
	_, err = db.Exec(sqlStr, info.Username, info.Password, info.Level)
	return err
}

func Login(db *sql.DB, info model.LoginInfo) (match bool, err error) {
	sqlStr := "SELECT EXISTS(SELECT 1 FROM user WHERE username = ? AND password = ?)"
	err = db.QueryRow(sqlStr, info.Username, info.Password).Scan(&match)
	return match, err
}

func GetID(db *sql.DB, username string) (id int64, err error) {
	sqlStr := "SELECT id FROM user WHERE username = ?"
	err = db.QueryRow(sqlStr, username).Scan(&id)
	return id, err
}

func GetLevel(db *sql.DB, username string) (level int8, err error) {
	sqlStr := "SELECT level FROM user WHERE username = ?"
	err = db.QueryRow(sqlStr, username).Scan(&level)
	return level, err
}

func GetUserInfo(db *sql.DB, userid int64) (info model.UserInfo, err error) {
	sqlStr := "SELECT username, level FROM user WHERE id = ?"
	err = db.QueryRow(sqlStr, userid).Scan(&info.Username, &info.Level)
	return info, err
}
