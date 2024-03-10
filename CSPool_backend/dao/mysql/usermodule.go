package mysqlmodule

import (
	"log"
	"main/model"
)

func CheckUserExist(username string) (exist bool) {
	sqlStr := "SELECT EXISTS(SELECT 1 FROM user WHERE username = ?)"
	err := Sdb.QueryRow(sqlStr, username).Scan(&exist)
	if err != nil {
		log.Fatalf("Failed to check user exist: %v", err)
	}
	return exist
}

func InvitationVerify(invitationcode string) int8 {
	var exists bool
	sqlStr := "SELECT EXISTS(SELECT 1 FROM invitation WHERE code = ?)"
	err := Sdb.QueryRow(sqlStr, invitationcode).Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check invitation code: %v", err)
	}
	if exists {
		return 2
	}
	return 3
}

func InsertUser(info model.UserInfo) {
	sqlStr := "insert into user(username, password, level) values(?,?,?)"
	_, err := Sdb.Exec(sqlStr, info.Username, info.Password, info.Level)
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}
	return
}

func Login(info model.LoginInfo) (match bool) {
	sqlStr := "SELECT EXISTS(SELECT 1 FROM user WHERE username = ? AND password = ?)"
	err := Sdb.QueryRow(sqlStr, info.Username, info.Password).Scan(&match)
	if err != nil {
		log.Fatalf("Failed to login: %v", err)
	}
	return match
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
