package main

import (
	"crawler/config"
	"crawler/engine"
	"crawler/persist"
	"crawler/redis"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {

	itemChan, err := persist.ItemSaver(config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	redis, err := redis.CreateConn(config.RedisPort)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
		RedisConn:        redis,
	}
	e.Run(engine.Request{
		URL:    config.EntryUrl,
		Parser: engine.CreateFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}
