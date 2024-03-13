package mysqlmodule

import (
	"database/sql"
	"errors"
	"main/model"
)

func CheckTitleExist(db *sql.DB, title string) (exist bool, err error) {
	sqlStr := "SELECT EXISTS(SELECT 1 FROM post WHERE title = ?)"
	err = db.QueryRow(sqlStr, title).Scan(&exist)
	return exist, err
}

func InsertPost(db *sql.DB, info model.PostInfo) (id int64, err error) {
	sqlStr := "insert into post(title,description,time,link,tag,authorname) values(?,?,?,?,?,?)"
	result, err := db.Exec(sqlStr, info.PostTitle, info.PostDescription, info.PostTime, info.PostLink, info.PostTag, info.AuthorName)
	if err != nil {
		return 0, err
	}
	// 获取最后插入行的ID
	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	// 返回最后插入行的ID
	return id, nil
}

func InsertUnderreview(db *sql.DB, id int64) (err error) {
	sqlStr := "insert into post_underreview(id) values(?)"
	_, err = db.Exec(sqlStr, id)
	return err
}

func DeleteUnderreview(db *sql.DB, id int64) (err error) {
	sqlStr := "DELETE FROM post_underreview WHERE id =?"
	_, err = db.Exec(sqlStr, id)
	return err
}
func AuthorizePost(db *sql.DB, vid int64) (err error) {
	sqlStr := "UPDATE post SET status = 1 WHERE id = ?"
	_, err = db.Exec(sqlStr, vid)
	return err
}

func CheckPidExist(db *sql.DB, vid int64) (exist bool, err error) {
	sqlStr := "SELECT EXISTS(SELECT 1 FROM post WHERE id = ?)"
	err = db.QueryRow(sqlStr, vid).Scan(&exist)
	return exist, err
}

func GetPostByID(db *sql.DB, vid int64) (info model.PostInfo, err error) {
	sqlStr := "SELECT * FROM post WHERE id = ?"
	err = db.QueryRow(sqlStr, vid).Scan(&info.IssueStatus, &info.PostID, &info.PostLike, &info.PostTitle, &info.PostDescription, &info.PostTime, &info.PostLink, &info.PostTag, &info.AuthorName)
	return info, err
}

func CheckPostStatus(db *sql.DB, vid int64) (status int, err error) {
	sqlStr := "SELECT status FROM post WHERE id = ?"
	err = db.QueryRow(sqlStr, vid).Scan(&status)
	return status, err
}

func LikePost(db *sql.DB, vid int64) (err error) {
	sqlStr := "UPDATE post SET `like` = `like` + 1  WHERE id = ?"
	_, err = db.Exec(sqlStr, vid)
	return err
}

func DislikePost(db *sql.DB, vid int64) (err error) {
	sqlStr := "UPDATE post SET `like` = `like` - 1  WHERE id = ?"
	_, err = db.Exec(sqlStr, vid)
	return err
}

func CheckVoteStatus(db *sql.DB, id int64, vid int64) (status int, err error) {
	sqlStr := "SELECT vote_type FROM post_vote WHERE user_id = ? AND video_id = ?"
	err = db.QueryRow(sqlStr, id, vid).Scan(&status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 5, nil
		} else {
			return 0, err
		}
	}
	return status, nil
}

func VoteInit(db *sql.DB, id int64, vid int64) (err error) {
	sqlStr := "insert into post_vote(user_id, video_id, vote_type) values(?,?,?)"
	_, err = db.Exec(sqlStr, id, vid, 0)
	return err
}

func UpdateVote(db *sql.DB, info model.VoteInfo, newstatus int) (err error) {
	sqlStr := "UPDATE post_vote SET `vote_type` = ?,vote_time = ?  WHERE user_id = ? AND video_id = ?"
	_, err = db.Exec(sqlStr, newstatus, info.VoteTime, info.VoterID, info.PostIDInt64)
	return err
}
