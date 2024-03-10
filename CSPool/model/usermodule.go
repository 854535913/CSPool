package model

type RegisterInfo struct {
	Username       string `json:"username" binding:"required"`
	Password       string `json:"password" binding:"required"`
	RePassword     string `json:"re_password" binding:"required,eqfield=Password"`
	InvitationCode string `json:"invitation"`
}

type LoginInfo struct {
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

type UserInfo struct {
	UserID   int64  `db:"userid"`
	Username string `db:"username" binding:"required"`
	Password string `db:"password" binding:"required"`
	Level    int8   `db:"level" binding:"required"`
}
