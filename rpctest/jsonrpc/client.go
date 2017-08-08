package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"time"
)

type Args struct{}

func main() {
	now := time.Now()
	client, err := jsonrpc.Dial("tcp", "192.168.2.83:8089")
	if err != nil {
		log.Fatal("dialing error:", err)
	}

	args := &Args{}
	reply := make([]string, 20)

	err = client.Call("Admin.Hello", args, &reply)
	if err != nil {
		log.Fatal("admin error:", err)
	}
	fmt.Println("jsonrpc :", time.Since(now).Seconds(), reply)
}
