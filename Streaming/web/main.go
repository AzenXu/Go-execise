package main

import (
	"net/http"
	"daker.wang/Azen/Go-execise/Streaming/web/server"
)

func main() {
	http.ListenAndServe(":8080", server.Regist())
}
