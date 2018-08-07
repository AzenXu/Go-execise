package routers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"io"
)

func Regist() (r *httprouter.Router) {
	r = httprouter.New()

	r.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		io.WriteString(writer, "哇又被你抓到啦~~羞羞羞👻")
	})

	return r
}