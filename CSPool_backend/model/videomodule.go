package model

import "time"

type VideoInfo struct {
	IssueStatus      bool      `db:"status"`
	VideoID          int64     `db:"id"`
	VideoLike        int64     `db:"like"`
	VideoTime        time.Time `db:"time"`
	VideoTitle       string    `json:"title" binding:"required" db:"title"`
	VideoDescription string    `json:"description" binding:"required" db:"description"`
	VideoLink        string    `json:"link" db:"link"`
	VideoTag         string    `json:"tag" db:"tag"`
	AuthorName       string    `db:"authorname"`
}

type VoteInfo struct {
	VoterID       int64
	VideoIDInt64  int64
	VideoIDString string
	VoteAction    string
	VoteTime      time.Time
}
