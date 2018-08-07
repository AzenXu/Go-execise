package defs

import (
	"net/http"
)

var (
	SessionTTLs = 60 * 30 // 30 min
	BucketCapacity  = 2
)

//  errors
var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC: http.StatusBadRequest, Error: Error{Content: "Request body is not correct", Code: "001"}}
	ErrorNotAuthUser = ErrorResponse{HttpSC: http.StatusUnauthorized, Error: Error{Content: "User authentication failed.", Code: "002"}}
	ErrorDBError = ErrorResponse{HttpSC: http.StatusInternalServerError, Error: Error{Content: "DB ops failed", Code: "003"}}
	ErrorInternalFaults = ErrorResponse{HttpSC: http.StatusInternalServerError, Error: Error{Content: "Internal service error", Code: "004"}}
	ErrorLimiterError = ErrorResponse{HttpSC: http.StatusTooManyRequests, Error: Error{Content: "Limiter failed", Code: "005"}}
)
