package auth

import (
	"net/http"
	"daker.wang/Azen/Go-execise/Streaming/api/session"
)

var HeaderFieldSession = "X-Session-Id"
var HeaderFieldUname = "X-User-Name"

func ValidateSession(request *http.Request) (bool) {
	sid := request.Header.Get(HeaderFieldSession)
	if len(sid) <= 0 {
		return false
	}

	if session.IsSessionExpired(sid) {
		return false
	}

	return true
}
