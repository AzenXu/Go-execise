package runnner

import (
	"time"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
)

type Worker struct {
	engine *Engine
	ticker *time.Ticker
}

func NewWorker(interval time.Duration, engine *Engine) *Worker {
	return &Worker{
		engine:engine,
		ticker:time.NewTicker(interval),
	}
}

func (w *Worker) startWork() {
	index := 0
	for {
		select {
		case <- w.ticker.C:
			log.Warn("🚀 第%d次定时器就要开动了~", index)
			index++
			go w.engine.StartAll()
		}
	}
}

func Start() {
	eng := New(3, VideoClearDispatcher, VideoClearExecutor)
	worker := NewWorker(30 * time.Second, eng)
	go worker.startWork()
}

func StartOld(duration time.Duration, task func()) {
	//  定时器怎么开？iOS里有NSTimer，Golang有木有类似的东东？查查文档
	//  文档中的time包，有两个type：Time、Ticker
	/*
	Timer代表单次时间事件，当Timer到期时，当时的时间会被发送给C，除非Timer是被AfterFunc函数创建的
	type Timer struct {
        C <-chan Time
        // 内含隐藏或非导出字段
	}

	Ticker是一个周期性传递时间的通道
	type Ticker struct {
        C <-chan Time // 周期性传递时间信息的通道
        // 内含隐藏或非导出字段
	}

	发现Tick的方法蛮贴切的
	*/
	ticker := time.NewTicker(duration)

	for {
		select {
		case <- ticker.C:
			fmt.Println("小白：时候到了..")
			task()
		}
	}
}
