package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

type UserResult struct {
	UserID string `json:"id"`
	OK bool `json:"status"`
}

type Error struct {
	Code string `json:"error_code"`
	Content string `json:"error_content"`
}

type ErrorResponse struct {
	Error Error
	HttpSC int
}

type Session struct {
	SessionID string `json:"session_id"`
	TTL int64 `json:"TTL"`
	UserName string `json:"login_name"`
}

type SessionResult struct {
	SessionID string `json:"session_id"`
	OK bool `string:"status"`
}

type NewVideo struct {
	AuthorId string `json:"author_id"`
	Name string `json:"name"`
}

type VideosInfo struct {
	Videos []*VideoInfo `json:"videos"`
}

type VideoInfo struct {
	Id string `json:"id"`
	AuthorId string `json:"author_id"`
	Name string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type CommentRequest struct {
	AuthorId string `json:"author_id"`
	Content string `json:"content"`
}

//  通过两张表inner join检索出来的
type VideoComment struct {
	Id string `json:"id"`
	VideoId string `json:"video_id"`
	Author string `json:"author"`
	Content string `json:"content"`
}

type Comments struct {
	Comments []*VideoComment `json:"comments"`
}
