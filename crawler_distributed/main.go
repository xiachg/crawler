package main

import (
	"../engine"

	"../scheduler"
	"../zhenai/parser"
	"./persist/client"
)

func main() {

	itemChan, err := client.ItemSaver(":1234")

	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
