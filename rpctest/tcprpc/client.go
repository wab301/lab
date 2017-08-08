package main

import (
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	address, err := net.ResolveTCPAddr("tcp", "192.168.2.83:8088")
	if err != nil {
		panic(err)
	}

	conn, _ := net.DialTCP("tcp", nil, address)
	defer conn.Close()

	client := rpc.NewClient(conn)
	defer client.Close()

	args := &Args{7, 8}
	reply := make([]string, 20)
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith errot :", err)
	}

	log.Fatalln(reply)
}
