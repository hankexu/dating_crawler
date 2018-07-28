package persist

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSave(t *testing.T) {
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
	const index = "dating_test"

	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	err = save(client, index, expected)
	if err != nil {
		t.Errorf("Save failed: %v", err)
	}

	if err != nil {
		panic(err)
	}
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}

}
