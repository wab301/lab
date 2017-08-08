package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "192.168.2.83:8088")
	if err != nil {
		log.Fatal("dailing :", err)
	}

	args := &Args{7, 8}
	reply := make([]string, 20)
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	log.Println(reply)
}
