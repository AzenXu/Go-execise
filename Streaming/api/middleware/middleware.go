package middleware

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
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

	//  鉴权

	m.r.ServeHTTP(w, re)
}



