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
		},
	}
	if actual != expected {
		t.Errorf("expected %v; but was %v", expected, actual)
	}
}
