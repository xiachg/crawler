package client

import (
	"../../../engine"
	"../../rpcsupport"
	"../../worker"
)

func CreateProcessor() (engine.Processor, error) {

	client, err := rpcsupport.NewClient(":9000")

	if err != nil {
		return nil, err
	}

	return func(req engine.Request) (engine.ParseResult, error) {

		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult

		err = client.Call("CrawlService.Process", sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}

		return worker.DeserializeResult(sResult), nil

	}, nil

}
