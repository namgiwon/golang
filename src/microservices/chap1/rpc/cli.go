package main

import (
	"fmt"

	"microservices/chap1/rpc/client"
	"microservices/chap1/rpc/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
}
