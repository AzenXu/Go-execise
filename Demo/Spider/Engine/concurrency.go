package engine

import (
	"daker.wang/Azen/Go-execise/Demo/Spider/Fetcher"
	"github.com/gpmgo/gopm/modules/log"
	"fmt"
)

type Concurrency struct {
	queue []Request
}

func (engine *Concurrency) Run(seeds ...Request) {
	engine.queue = append(engine.queue, seeds...)

	for len(engine.queue) > 0  {

		r := engine.queue[0]
		engine.queue = engine.queue[1:]

		items := worker(r)

		for i :=0; i < len(items); i++  {
			item := items[i]
			fmt.Println(item.Name)
			engine.queue = append(engine.queue, item.Request)
		}
	}
}

func worker(r Request) (items []Item) {
	result, err := fetcher.Fetch(r.URL)

	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	return r.ParasFunc(result)
}