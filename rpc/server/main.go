package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"../../rpc"
)

func main() {
	// 注册rpc消息
	rpc.Register(rpcdemo.DemoService{})
	// 开启rpc服务
	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	for {

		conn, err := listener.Accept()

		if err != nil {
			log.Printf("accpect error: %v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}

}
