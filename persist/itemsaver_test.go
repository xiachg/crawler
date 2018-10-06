package persist

import (
	"context"
	"encoding/json"
	"testing"

	"gopkg.in/olivere/elastic.v5"

	"../model"

	"../engine"
)

func TestSave(t *testing.T) {

	expected := engine.Item{
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

	// TODO: Try to start up elastic search here using docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	const index string = "dating_test"

	err = Save(client, index, expected)

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", *resp.Source)

	var actual engine.Item

	json.Unmarshal(*resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)

	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}

}
