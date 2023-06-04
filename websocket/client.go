package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()

	// 发送消息
	for true {
		var s string
		fmt.Scan(&s)
		err = conn.WriteMessage(websocket.TextMessage, []byte(s))
		if err != nil {
			log.Fatal(err)
		}
		a := time.Now().Unix()
		//t := time.NewTicker(time.Second * 5)
		// 读取消息
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		b := time.Now().Unix()
		if (b - a) > 3 {
			break
		} else {
			log.Println("Received message:", string(p))
		}
	}
}
