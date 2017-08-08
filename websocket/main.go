package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	var err error
	var i int
	for {
		time.AfterFunc(2*time.Second, func() {
			websocket.Message.Send(ws, "this is a test")
		})
	}
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + strconv.Itoa(i) + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
