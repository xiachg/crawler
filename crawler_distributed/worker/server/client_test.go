package main

import (
	"fmt"
	"testing"
	"time"

	"../../rpcsupport"
	"../../worker"

	"../../../config"
)

func TestCrawlService(t *testing.T) {

	const host = ":9000"

	go rpcsupport.ServeRpc(host, worker.CrawlService{})

	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)

	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1029982807",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "Lucy",
		},
	}

	var result worker.ParseResult

	err = client.Call("CrawlService.Process", req, &result)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}

}
