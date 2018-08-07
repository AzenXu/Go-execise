package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
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
