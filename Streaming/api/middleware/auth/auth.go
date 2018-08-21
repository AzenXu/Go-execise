package auth

import (
	"net/http"
	"daker.wang/Azen/Go-execise/Streaming/api/session"
	"github.com/gpmgo/gopm/modules/log"
)

var HeaderFieldSession = "X-Session-Id"
var HeaderFieldUname = "X-User-Name"

func ValidateSession(request *http.Request) (bool) {
	sid := request.Header.Get(HeaderFieldSession)
	log.Warn("👏 抓到sid:%s", sid)
	if len(sid) <= 0 {
		return false
	}

	if !session.IsSessionUseful(sid) {
		return false
	}

	return true
}
