package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

type Args struct {
	A, B int
}

type Arith int

func (*Arith) Multiply(args *Args, reply *([]string)) error {
	*reply = append(*reply, fmt.Sprintf("A * B = %v", args.A*args.B))
	return nil
}

func main() {
	newServer := rpc.NewServer()
	newServer.Register(new(Arith))

	l, err := net.Listen("tcp", "192.168.2.83:8088")
	if err != nil {
		log.Fatalf("net:Listen tcp :0:%v", err)
	}

	go newServer.Accept(l)
	newServer.HandleHTTP("/foo", "/bar")
	time.Sleep(10 * time.Second)
}
