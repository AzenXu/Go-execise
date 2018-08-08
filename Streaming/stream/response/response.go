package response

import (
	"net/http"
	"io"
	"daker.wang/Azen/Go-execise/Streaming/api/defs"
)

func SendErrorResponse(w http.ResponseWriter, err defs.ErrorResponse) {
	w.WriteHeader(err.HttpSC)
	content := "我是流媒体服务器，我错了！我再也不敢了~~ 555~~~" + err.Error.Content
	io.WriteString(w, content)
}

func SendNormalResponse(w http.ResponseWriter, content string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, content)
}
