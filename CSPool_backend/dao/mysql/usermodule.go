package mysqlmodule

import (
	"database/sql"
	"log"
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

func Login(info model.LoginInfo) (match bool, err error) {
	sqlStr := "SELECT EXISTS(SELECT 1 FROM user WHERE username = ? AND password = ?)"
	err = Sdb.QueryRow(sqlStr, info.Username, info.Password).Scan(&match)
	return match, err
}

func GetID(username string) (id int64) {
	sqlStr := "SELECT id FROM user WHERE username = ?"
	err := Sdb.QueryRow(sqlStr, username).Scan(&id)
	if err != nil {
		log.Fatalf("Failed to get ID: %v", err)
	}
	return id
}

func GetLevel(username string) (level int8) {
	sqlStr := "SELECT level FROM user WHERE username = ?"
	err := Sdb.QueryRow(sqlStr, username).Scan(&level)
	if err != nil {
		log.Fatalf("Failed to get level: %v", err)
	}
	return level
}
