package main

import (
	"log"

	"../../rpcsupport"
	"../../worker"
)

func main() {

	log.Fatal(rpcsupport.ServeRpc(":9000", worker.CrawlService{}))

}
