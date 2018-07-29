package parser_test

import (
	"crawler/engine"
	"crawler/model"
	"crawler/zhenai/parser"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("./profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := parser.ParseProfile(contents, "http://album.zhenai.com/u/110171680", "幽诺")

	if len(result.Items) != 1 {
		t.Errorf("result should have %d requests;"+
			"but had %d", 1, len(result.Items))
	}

	actual := result.Items[0]
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/110171680",
		Type: "zhenai",
		Id:   "110171680",
		Payload: model.Profile{
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
	if actual != expected {
		t.Errorf("expected %v;\n but was %v", expected, actual)
	}

	//for i, url := range expectedUrls {
	//	if result.Requests[i].URL!=url {
	//		t.Errorf("expected url #%d: %s; but was %d", i,url, result.Requests[i].URL)
	//	}
	//}
	//if len(result.Items) !=resultSize {
	//	t.Errorf("result should have %d items;" +
	//		"but had %d", resultSize, len(result.Requests))
	//}
	//
	//for i, city := range expectedCities {
	//	if result.Items[i].(string)!= city {
	//		t.Errorf("expected city #%d: %s; but was %s", i,city, result.Items[i].(string))
	//	}
	//}

}
