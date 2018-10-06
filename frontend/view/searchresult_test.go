package view

import (
	"os"
	"testing"

	"../model"

	"../../engine"

	common "../../model"
)

func TestSearchResultView_Render(t *testing.T) {

	view := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1029982807",
		Type: "zhenai",
		Id:   "1029982807",
		Payload: common.Profile{
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

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)

	if err != nil {
		panic(err)
	}

}
