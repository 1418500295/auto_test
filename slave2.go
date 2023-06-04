package main

import (
	"auto_test/utils"
	"log"
	"net"
	"net/rpc"
)

type Slave2Service struct { //空结构体

}

func (s *Slave2Service) Slave2(request string, reply *string) error { //为hello结构体绑定hello方法
	*reply = utils.Do() //拼接字符串
	return nil          //返回空值

}

func main() {
	rpc.RegisterName("Slave2Service", new(Slave2Service))
	listener, err := net.Listen("tcp", ":8881")
	if err != nil {
		log.Fatal("ListenTCP erro :", err)

	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	rpc.ServeConn(conn)
}
