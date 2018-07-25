package main

import (
	"daker.wang/Azen/Go-execise/Demo/Spider/Engine"
	"daker.wang/Azen/Go-execise/Demo/Spider/zhenai/Parser"
)

func main() {
	var simpleEngine = new(engine.SimpleEngine)
	simpleEngine.Run(engine.Request{
		URL:"http://www.zhenai.com/zhenghun",
		ParasFunc: parser.PickUpCitys,
	})
}
