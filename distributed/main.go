package main

import (
	"crawler/config"
	"crawler/engine"
	"crawler/model"
	"crawler/persist"
	"crawler/persist/client"
	"crawler/redis"
	"crawler/rpcsupport"
	"crawler/scheduler"
	worker "crawler/worker/client"
	"crawler/zhenai/parser"
	"flag"
	"log"
	"net/rpc"
)

func main() {
	flag.Parse()
	itemChan, err := persist.ItemSaver(config.ElasticIndex)
	if err != nil {
		panic(err)
	}
	conf, err := model.LoadConfig("config.json")
	if err != nil {
		panic(err)
	}
	pool := createClientPool(conf.Workers)
	processor := worker.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}
	itemChan, err = client.ItemSaver(conf.Saver)
	if err != nil {
		panic(err)
	}

	red, err := redis.CreateConn(config.RedisPort)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
		RedisConn:        red,
	}
	e.Run(engine.Request{
		URL:    "http://www.zhenai.com/zhenghun",
		Parser: engine.CreateFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err != nil {
			log.Printf("error connecting to %s: %v", h, err)
		} else {
			log.Printf("Connected to %s", h)
			clients = append(clients, client)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
