package main

import (
	"net/http"
	"daker.wang/Azen/Go-execise/Streaming/dispatcher/router"
	"daker.wang/Azen/Go-execise/Streaming/dispatcher/runnner"
)

func main() {
	//  开启一个定时器，定时执行一个任务
	runnner.Start()
	http.ListenAndServe(":9002", router.Register())
}
