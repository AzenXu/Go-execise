package main

import (
	"net/http"
	"daker.wang/Azen/Go-execise/Streaming/api/routers"
	"daker.wang/Azen/Go-execise/Streaming/api/middleware"
)

func main() {
	m := middleware.New(routers.Register())
	http.ListenAndServe(":9000", m)
}
