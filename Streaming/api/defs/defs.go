package defs

import (
	"net/http"
)

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
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC: http.StatusBadRequest, Error: Error{Content: "Request body is not correct", Code: "001"}}
	ErrorNotAuthUser = ErrorResponse{HttpSC: http.StatusUnauthorized, Error: Error{Content: "User authentication failed.", Code: "002"}}
	ErrorDBError = ErrorResponse{HttpSC: http.StatusInternalServerError, Error: Error{Content: "DB ops failed", Code: "003"}}
	ErrorInternalFaults = ErrorResponse{HttpSC: http.StatusInternalServerError, Error: Error{Content: "Internal service error", Code: "004"}}
)
