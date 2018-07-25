package main

import (
	"daker.wang/Azen/Go-execise/Demo/Spider/Engine"
	"daker.wang/Azen/Go-execise/Demo/Spider/zhenai/Parser"
	"daker.wang/Azen/Go-execise/Demo/Spider/Scheduler"
)

func main() {
	//runSimpleEngine()
	runConcurrencyEngine()
}

func runConcurrencyEngine() {
	var concurrencyEngine = engine.Concurrency{
		Scheduler:&scheduler.SimpleScheduler{
			RequestChannel:make(chan engine.Request),
		},
		WorkerCount:10,
	}
	concurrencyEngine.Run(engine.Request{
		URL:"http://www.zhenai.com/zhenghun",
		ParasFunc: parser.PickUpCitys,
	})
}

func runSimpleEngine() {
	var simpleEngine = new(engine.SimpleEngine)
	simpleEngine.Run(engine.Request{
		URL:"http://www.zhenai.com/zhenghun",
		ParasFunc: parser.PickUpCitys,
	})
}
