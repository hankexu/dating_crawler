package main

import (
	"crawler/rpcsupport"
	"crawler/worker"
	"flag"
	"fmt"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{})
}