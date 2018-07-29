package view

import (
	"crawler/engine"
	"crawler/frontend/model"
	common "crawler/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")
	page := model.SearchResult{}
	page.Hits = 100
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/110171680",
		Type: "zhenai",
		Id:   "110171680",
		Payload: common.Profile{
			Name:          "幽诺",
			Age:           25,
			Height:        160,
			Weight:        0,
			Income:        "3000元以下",
			Education:     "中专",
			Marriage:      "离异",
			Gender:        "女",
			Constellation: "处女座",
			Car:           "未购车",
			Address:       "四川阿坝",
			Photo:         "http://photo15.zastatic.com/images/photo/27543/110171680/1515042497285.jpg?scrop=1&crop=1&w=650&h=650&cpos=north",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}
	out, err := os.Create("search_result_test.html")
	if err != nil {
		panic(err)
	}
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}

}
