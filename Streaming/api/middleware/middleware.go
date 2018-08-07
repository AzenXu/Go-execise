package middleware

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"daker.wang/Azen/Go-execise/Streaming/api/middleware/auth"
)

type MiddleWare struct {
	r *httprouter.Router
}

func New(router *httprouter.Router) *MiddleWare {
	m := &MiddleWare{}
	m.r = router
	return m
}

func (m *MiddleWare) ServeHTTP(w http.ResponseWriter, re *http.Request) {

	result := auth.ValidateSession(re)
	log.Println("鉴权：结果为：", result)

	m.r.ServeHTTP(w, re)
}





