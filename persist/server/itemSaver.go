package main

import (
	"crawler/config"
	"crawler/persist"
	"crawler/rpcsupport"
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serverRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serverRpc(host string, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{Client: client, Index: index})

	return nil
}
