package main

import (
	"testing"
	"time"

	"../../../engine"
	"../../../model"
	"../../rpcsupport"
)

func TestItemSaver(t *testing.T) {

	const host = ":1234"

	// start ItemSaverServer
	go serveRpc(host, "test_rpc")

	time.Sleep(time.Second)

	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)

	if err != nil {
		panic(err)
	}

	// item data
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1029982807",
		Type: "zhenai",
		Id:   "1029982807",
		Payload: model.Profile{
			Name:       "Lucy",
			Gender:     "女",
			Age:        22,
			Height:     170,
			Weight:     49,
			Income:     "8001-12000元",
			Marriage:   "未婚",
			Education:  "大学本科",
			Occupation: "财务/申计",
			Hokou:      "上海浦东新区",
			Xinzuo:     "狮子座",
			House:      "和家人同住",
			Car:        "未购车",
		},
	}

	// Call Save
	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; error: %s", result, err)
	}

}
