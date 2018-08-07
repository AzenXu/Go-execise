package response

import (
	"net/http"
	"io"
	"daker.wang/Azen/Go-execise/Streaming/api/defs"
)

func SendErrorResponse(w http.ResponseWriter, err defs.ErrorResponse) {
	w.WriteHeader(err.HttpSC)
	content := "我错了！！！555~~~" + err.Error.Content
	io.WriteString(w, content)
}

func SendNormalResponse(w http.ResponseWriter, content string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, content)
}
