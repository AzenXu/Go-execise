package main

import (
	"daker.wang/Azen/Go-execise/Demo/Spider/Engine"
	"daker.wang/Azen/Go-execise/Demo/Spider/zhenai/Parser"
	"daker.wang/Azen/Go-execise/Demo/Spider/Scheduler"
)

func main() {
	//runSimpleEngine()
	//runConcurrencyEngine()
	customFetchTest()
	//runQueueEngine()
}

func runQueueEngine() {
	queueEngine := engine.Queue{
		Scheduler:&scheduler.QueueScheduler{},
		WorkerCount:10,
	}
	queueEngine.Run(engine.Request{
		URL:"http://www.zhenai.com/zhenghun",
		ParasFunc: parser.PickUpCitys,
	})
}

func customFetchTest() {
	var concurrencyEngine = engine.Concurrency{
		Scheduler:&scheduler.SimpleScheduler{
			RequestChannel:make(chan engine.Request),
		},
		WorkerCount:10,
	}

	concurrencyEngine.Run(engine.Request{
		URL:"http://www.zhenai.com/zhenghun/changzhi",
		ParasFunc: func(bytes []byte) []engine.Item {
			return parser.PickUpPersons(bytes)
		},
	})
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
