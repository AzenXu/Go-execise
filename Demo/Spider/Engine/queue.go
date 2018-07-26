package engine

import (
	"daker.wang/Azen/Go-execise/Demo/Spider/Fetcher"
	"github.com/gpmgo/gopm/modules/log"
)

type QueueScheduler interface {
	Submit(request Request)
	GetRequsetChannel() (requestChannel chan Request)
	SubmitWorker(worker chan Request)
	Run()
}

type Queue struct {
	Scheduler QueueScheduler
	WorkerCount int
}

func (engine *Queue) Run(seeds ...Request) {

	requestChannel := engine.Scheduler.GetRequsetChannel()
	resultsOutChannel := make(chan []Item)

	engine.Scheduler.Run()

	for i := 0; i < engine.WorkerCount; i++ {
		createQueueWork(requestChannel, resultsOutChannel, engine.Scheduler)
	}

	for i := 0; i < len(seeds); i++ {
		r := seeds[i]
		engine.Scheduler.Submit(r)
	}

	for {
		results := <- resultsOutChannel
		for _, result := range results {
			if len(result.Request.URL) <= 0 { // 不是一个有效URL则抛弃
				continue
			}
			engine.Scheduler.Submit(result.Request)
		}
	}
}

func createQueueWork(requestChannel chan Request, itemsOut chan []Item, scheduler QueueScheduler) {
	go func() {
		for {
			// 一旦空闲，先submitWorker
			scheduler.SubmitWorker(requestChannel)

			// 等待分配到任务
			r := <-requestChannel
			itemsOut <- queueWorker(r)
		}
	}()
}

func queueWorker(r Request) (items []Item) {
	result, err := fetcher.Fetch(r.URL)

	if err != nil {
		log.Error(err.Error())
	}

	return r.ParasFunc(result)
}