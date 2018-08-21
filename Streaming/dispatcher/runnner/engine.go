package runnner

import (
	"log"
)

type fn func(data chan string) *TaskResult

type Engine struct {
	Data chan string // 待删数据
	DataSize int
	Control chan int // 控制 1 生产完毕 2 消费完毕
	Err chan int // 一有信息马上终止
	Dispatcher fn
	Executor fn
}

func New(size int, d fn, e fn) (en *Engine) {
	return &Engine{
		Data:make(chan string, size),
		Control:make(chan int, 1),
		Err:make(chan int, 1),
		DataSize:size,
		Dispatcher:d,
		Executor:e,
	}
}

func(e *Engine) startDispatch() {
	for {
		select {
		case <- e.Err:
			//  终止
			return
		case c := <- e.Control:
			if c == 1 {
				//  开始消费
				log.Println("开始消费")
				tr := e.Executor(e.Data)
				if tr.Err != nil || tr.Event == TaskEventClose {
					e.Err <- 0
				} else {
					e.Control <- 2
				}
			} else if c == 2 {
				//  开始生产
				log.Println("开始生产")
				tr := e.Dispatcher(e.Data)
				if tr.Err != nil || tr.Event == TaskEventClose {
					e.Err <- 0
					log.Println("💋 本次生产任务已经全部结束！")
				} else {
					e.Control <- 1
				}
			}
		}
	}
}

func(e *Engine) StartAll() {
	e.Control <- 2
	e.startDispatch()
}
