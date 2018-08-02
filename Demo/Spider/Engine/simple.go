package engine

import (
	"daker.wang/Azen/Go-execise/Demo/Spider/Fetcher"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
)

type SimpleEngine struct {
	queue []Request
}

func (engine *SimpleEngine) Run(seeds ...Request) {
	engine.queue = append(engine.queue, seeds...)

	for len(engine.queue) > 0 {

		r := engine.queue[0]
		engine.queue = engine.queue[1:]

		result, err := fetcher.Fetch(r.URL)

		if err != nil {
			log.Error(err.Error())
		}

		items := r.ParasFunc(result)

		for i := 0; i < len(items); i++ {
			item := items[i]
			fmt.Println(item.Name)
			engine.queue = append(engine.queue, item.Request)
		}
	}
}
