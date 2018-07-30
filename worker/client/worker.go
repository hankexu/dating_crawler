package client

import (
	"crawler/config"
	"crawler/engine"
	"crawler/worker"
	"net/rpc"
)

func CreateProcessor(clientsChan chan *rpc.Client) engine.Processor {
	return func(r engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(r)
		var sResult worker.ParseResult
		c := <-clientsChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
