package main

import (
	"net/http"
	"daker.wang/Azen/Go-execise/Streaming/stream/middleware"
	"daker.wang/Azen/Go-execise/Streaming/stream/routers"
)

func main() {
	m := middleware.New(routers.Regist())
	http.ListenAndServe(":9001", m)
}
