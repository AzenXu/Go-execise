package main

import (
	"net/http"
	"daker.wang/Azen/Go-execise/Streaming/dispatcher/router"
	"daker.wang/Azen/Go-execise/Streaming/dispatcher/runnner"
	"time"
	"daker.wang/Azen/Go-execise/Streaming/dispatcher/ops"
)

func main() {

	//  开启一个定时器，定时执行一个任务
	runnner.Start(3 * time.Second, cleanHouse)

	http.ListenAndServe(":9002", router.Register())
}

func cleanHouse() {
	//  获取待清理文件列表
	vids := ops.LoadAllPaddingVIDs()
	//  清理文件
	for _, vid := range vids {
		ops.DeleteVideo(vid)
	}
}
