package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

type Args struct {
}

type Admin int

func (*Admin) Hello(args *Args, reply *([]string)) error {
	*reply = append(*reply, "Hello World!")
	return nil
}

func main() {
	HttpRPC()
	// JsonRPC()
	time.Sleep(90 * time.Second)
}

func HttpRPC() {
	admin := new(Admin)

	rpc.Register(admin)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatal("listen :", err)
	}
	go http.Serve(l, nil)
}

func JsonRPC() {
	admin := new(Admin)
	rpc.Register(admin)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8089")
	if err != nil {
		log.Fatal("tcpaddr :", err)
	}

	listen, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal("listen :", err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("accept :", err)
		}
		jsonrpc.ServeConn(conn)
	}
}
