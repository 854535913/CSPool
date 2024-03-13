package mysqlmodule

import "errors"

var (
	ErrorUserExist         = errors.New("用户已存在")
	ErrorUserNotExist      = errors.New("用户不存在")
	ErrorInvalidPassword   = errors.New("用户名或密码错误")
	ErrorInvalidID         = errors.New("无效的ID")
	ErrorTitleExist        = errors.New("标题已存在")
	ErrorInvalidPostID     = errors.New("无效的PostID")
	ErrorPostPublished     = errors.New("贴子已发表")
	ErrorVoteStatusLike    = errors.New("已经喜欢这个贴子")
	ErrorVoteStatusDislike = errors.New("已经不喜欢这个贴子")
	ErrorVoteStatusNone    = errors.New("已经取消了这个贴子")
	ErrorVoteStatus        = errors.New("无效的喜欢/不喜欢状态码")
	ErrorInvalidVoteAction = errors.New("无效的动作码")
)
