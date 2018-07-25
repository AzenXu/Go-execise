package scheduler

import (
	"daker.wang/Azen/Go-execise/Demo/Spider/Engine"
)

type SimpleScheduler struct {
	RequestChannel chan engine.Request
}

func (scheduler *SimpleScheduler) Submit(request engine.Request) {
	scheduler.RequestChannel <- request
}

func (scheduler *SimpleScheduler) GetRequsetChannel() (requestChannel chan engine.Request) {
	return scheduler.RequestChannel
}

