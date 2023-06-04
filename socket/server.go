package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err5 := net.Listen("tcp", "127.0.0.1:8888")
	if err5 != nil {
		fmt.Println(err5)
	}
	fmt.Println("服务端等待客户端连接中...")
	buf := make([]byte, 1024)
	for {
		con, err5 := listen.Accept()
		if err5 != nil {
			fmt.Println(err5)
		}
		//Do()
		n, err5 := con.Read(buf)
		if err5 != nil {
			fmt.Println(err5)
		}
		fmt.Println(string(buf[:n]))
		var in string
		_, _ = fmt.Scan(&in)

		con.Write([]byte(in))
	}
}
