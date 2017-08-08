package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct {
	A, B int
}

type Arith int

func (*Arith) Multiply(args *Args, reply *([]string)) error {
	*reply = append(*reply, "hello world! this is a test!!!!")
	*reply = append(*reply, fmt.Sprintf("A is %v,B is %v", args.A, args.B))
	return nil
}

func main() {
	arith := new(Arith)

	rpc.Register(arith)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	go http.Serve(l, nil)

	time.Sleep(5 * time.Second)
}
