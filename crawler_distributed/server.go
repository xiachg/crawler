package main

import (
	"../engine"

	worker "../crawler_distributed/worker/client"
	"../scheduler"
	"../zhenai/parser"
	"./persist/client"
)

func main() {

	itemChan, err := client.ItemSaver(":1234")

	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()

	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})

}
