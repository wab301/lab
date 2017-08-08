package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
}

type Admin int

func (*Admin) Hello(arg *Args, reply *([]string)) error {
	*reply = append(*reply, "hello world!")
	return nil
}

func main() {
	admin := new(Admin)
	rpc.Register(admin)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8089")
	if err != nil {
		log.Fatal("tcp error", err)
	}

	listen, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal("listen error", err)

	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}
