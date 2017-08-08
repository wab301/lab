package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

type Args struct {
}

func BenchmarkHTTPRPC(b *testing.B) {
	client, err := rpc.DialHTTP("tcp", "192.168.2.83:8088")
	if err != nil {
		b.Fatal("dialing :", err)
	}

	defer client.Close()

	args := &Args{}
	reply := make([]string, 20)

	client.Call("Admin.Hello", args, &reply)
	for i := 0; i < b.N; i++ {

		client.Call("Admin.Hello", args, &reply)
	}
	/*err = client.Call("Admin.Hello", args, &reply)
	if err != nil {
		t.Errorf("Call :", err)
	}*/
}

func BenchmarkJSONRPC(b *testing.B) {
	client, err := jsonrpc.Dial("tcp", "192.168.2.83:8089")
	if err != nil {
		b.Fatal("dialing :", err)
	}
	defer client.Close()

	args := &Args{}
	reply := make([]string, 20)

	for i := 0; i < b.N; i++ {
		client.Call("Admin.Hello", args, &reply)
	}

	// err = client.Call("Admin.Hello", args, &reply)
	// if err != nil {
	// 	t.Errorf("Call :", err)
	// }
}
