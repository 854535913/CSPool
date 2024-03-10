package mysqlmodule

import (
	"database/sql"
	"errors"
	"log"
	"main/model"
)

func CheckTitleExist(title string) (exist bool) {
	sqlStr := "SELECT EXISTS(SELECT 1 FROM video WHERE title = ?)"
	err := Sdb.QueryRow(sqlStr, title).Scan(&exist)
	if err != nil {
		log.Fatalf("Failed to check title exist: %v", err)
	}
	return exist
}

func InsertVideo(info model.VideoInfo) (int64, error) {
	// 插入数据的SQL语句
	sqlStr := "insert into video(title,description,time,link,tag,authorname) values(?,?,?,?,?,?)"
	// 执行SQL语句
	result, err := Sdb.Exec(sqlStr, info.VideoTitle, info.VideoDescription, info.VideoTime, info.VideoLink, info.VideoTag, info.AuthorName)
	if err != nil {
		return 0, err
	}
	// 获取最后插入行的ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	// 返回最后插入行的ID
	return id, nil
}

func InsertUnderreview(id int64) (err error) {
	// 插入数据的SQL语句
	sqlStr := "insert into video_underreview(id) values(?)"
	// 执行SQL语句
	_, err = Sdb.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return
}

func DeleteUnderreview(id int64) (err error) {
	// 插入数据的SQL语句
	sqlStr := "DELETE FROM video_underreview WHERE id =?"
	// 执行SQL语句
	_, err = Sdb.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return
}
func AuthorizeVideo(vid int64) {
	sqlStr := "UPDATE video SET status = 1 WHERE id = ?"
	_, err := Sdb.Exec(sqlStr, vid)
	if err != nil {
		log.Fatalf("Failed to authorize: %v", err)
	}
}

func CheckVidExist(vid int64) (exist bool) {
	sqlStr := "SELECT EXISTS(SELECT 1 FROM video WHERE id = ?)"
	err := Sdb.QueryRow(sqlStr, vid).Scan(&exist)
	if err != nil {
		log.Fatalf("Failed to check id exist: %v", err)
	}
	return exist
}

func GetVideoByID(vid int64) (info model.VideoInfo, err error) {
	sqlStr := "SELECT * FROM video WHERE id = ?"
	err = Sdb.QueryRow(sqlStr, vid).Scan(&info.IssueStatus, &info.VideoID, &info.VideoLike, &info.VideoTitle, &info.VideoDescription, &info.VideoTime, &info.VideoLink, &info.VideoTag, &info.AuthorName)
	if err != nil {
		return model.VideoInfo{}, err
	}
	return info, nil
}

func CheckVideoStatus(vid int64) (status int) {
	sqlStr := "SELECT status FROM video WHERE id = ?"
	err := Sdb.QueryRow(sqlStr, vid).Scan(&status)
	if err != nil {
		log.Fatalf("Failed to check vid status: %v", err)
	}
	return status
}

func LikeVideo(vid int64) (err error) {
	sqlStr := "UPDATE video SET `like` = `like` + 1  WHERE id = ?"
	_, err = Sdb.Exec(sqlStr, vid)
	if err != nil {
		log.Fatalf("Failed to change like: %v", err)
		return err
	}
	return
}

func DislikeVideo(vid int64) (err error) {
	sqlStr := "UPDATE video SET `like` = `like` - 1  WHERE id = ?"
	_, err = Sdb.Exec(sqlStr, vid)
	if err != nil {
		log.Fatalf("Failed to change like: %v", err)
		return err
	}
	return
}

func CheckVoteStatus(id int64, vid int64) (status int, err error) {
	sqlStr := "SELECT vote_type FROM video_vote WHERE user_id = ? AND video_id = ?"
	err = Sdb.QueryRow(sqlStr, id, vid).Scan(&status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 5, nil
		} else {
			log.Fatalf("Failed to check id vote: %v", err)
			return 0, err
		}
	}
	return status, nil
}

func VoteInit(id int64, vid int64) (err error) {
	sqlStr := "insert into video_vote(user_id, video_id, vote_type) values(?,?,?)"
	_, err = Sdb.Exec(sqlStr, id, vid, 0)
	if err != nil {
		log.Fatalf("Failed init vote: %v", err)
		return err
	}
	return
}

func UpdateVote(info model.VoteInfo, newstatus int) (err error) {
	sqlStr := "UPDATE video_vote SET `vote_type` = ?,vote_time = ?  WHERE user_id = ? AND video_id = ?"
	_, err = Sdb.Exec(sqlStr, newstatus, info.VoteTime, info.VoterID, info.VideoIDInt64)
	if err != nil {
		log.Fatalf("Failed to update vote info: %v", err)
		return err
	}
	return
}
