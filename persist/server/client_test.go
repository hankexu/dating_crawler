package main

import (
	"crawler/engine"
	"crawler/model"
	"crawler/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	go serverRpc(host, "test1")
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)

	}
	result := ""
	item := engine.Item{
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
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "OK" {
		t.Errorf("result: %s,err:%s ", result, err)
	}
}
