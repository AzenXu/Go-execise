package engine

import (
	"daker.wang/Azen/Go-execise/Demo/Spider/Fetcher"
	"github.com/gpmgo/gopm/modules/log"
)

type Scheduler interface {
	Submit(request Request)
	GetRequsetChannel() (requestChannel chan Request)
}

type Concurrency struct {
	Scheduler Scheduler
	WorkerCount int
}

func (engine *Concurrency) Run(seeds ...Request) {

	requestChannel := engine.Scheduler.GetRequsetChannel()
	resultsOutChannel := make(chan []Item)

	for i := 0; i < engine.WorkerCount; i++ {
		createWork(requestChannel, resultsOutChannel)
	}

	for i := 0; i < len(seeds); i++ {
		r := seeds[i]
		engine.Scheduler.Submit(r)
	}

	for {
		results := <- resultsOutChannel
		for _, result := range results {
			engine.Scheduler.Submit(result.Request)
		}
	}
}

func createWork(requestChannel chan Request, itemsOut chan []Item) {
	go func() {
		for {
			r := <-requestChannel
			itemsOut <- worker(r)
		}
	}()
}

func worker(r Request) (items []Item) {
	result, err := fetcher.Fetch(r.URL)

	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	return r.ParasFunc(result)
}