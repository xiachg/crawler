package main

import (
	"../../persist"
	"../../rpcsupport"
	"gopkg.in/olivere/elastic.v5"
)

func main() {

	// log.Fatal(serveRpc(":1234", "dating_profile"))

	err := serveRpc(":1234", "dating_profile")

	if err != nil {
		panic(err)
	}

}

func serveRpc(host, index string) error {

	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		return err
	}

	rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})

	return nil

}
