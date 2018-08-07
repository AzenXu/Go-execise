package middleware

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"daker.wang/Azen/Go-execise/Streaming/stream/middleware/limiter"
	"daker.wang/Azen/Go-execise/Streaming/api/defs"
	"daker.wang/Azen/Go-execise/Streaming/stream/response"
)

type MiddleWare struct {
	r *httprouter.Router
	l *limiter.Limiter
}

func New(router *httprouter.Router) *MiddleWare {
	m := &MiddleWare{}
	m.l = limiter.New(defs.BucketCapacity)
	m.r = router
	return m
}

func (m *MiddleWare) ServeHTTP(w http.ResponseWriter, re *http.Request) {

	if !m.l.GetToken() {
		response.SendErrorResponse(w, defs.ErrorLimiterError)
		return
	}

	m.r.ServeHTTP(w, re)

	m.l.ReleaseToken() // TODO: 测试此处加defer是否有错
}