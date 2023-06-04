package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err5 := net.Dial("tcp", "127.0.0.1:8888")
	if err5 != nil {
		log.Fatal(err5)
	}
	buf := make([]byte, 1024)

	fmt.Println("连接服务器成功。。。")
	for {
		var i string
		_, _ = fmt.Scan(&i)
		_, err5 = conn.Write([]byte(i))
		if err5 != nil {
			fmt.Println(err5)
		}
		n, err5 := conn.Read(buf)
		if err5 != nil {
			fmt.Println(err5)
		}
		fmt.Println(string(buf[:n]))
	}
}
