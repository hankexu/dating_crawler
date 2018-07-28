package parser_test

import (
	"crawler/zhenai/parser"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	const resultSize = 470

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	result := parser.ParseCityList(contents, "")

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests;"+
			"but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].URL != url {
			t.Errorf("expected url #%d: %s; but was %d", i, url, result.Requests[i].URL)
		}
	}
}
