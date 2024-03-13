package model

import "time"

type PostInfo struct {
	IssueStatus     bool      `db:"status"`
	PostID          int64     `db:"id"`
	PostLike        int64     `db:"like"`
	PostTime        time.Time `db:"time"`
	PostTitle       string    `json:"title" binding:"required" db:"title"`
	PostDescription string    `json:"description" binding:"required" db:"description"`
	PostLink        string    `json:"link" db:"link"`
	PostTag         string    `json:"tag" db:"tag"`
	AuthorName      string    `db:"authorname"`
}

type VoteInfo struct {
	VoterID      int64
	PostIDInt64  int64
	PostIDString string
	VoteAction   string
	VoteTime     time.Time
}
