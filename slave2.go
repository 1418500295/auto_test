package main

import (
	"auto_test/utils"
	"encoding/json"
	"fmt"
	"net"
)

//type Slave2Service struct { //空结构体
//
//}
//
//func (s *Slave2Service) Slave2(request string, reply *string) error { //为hello结构体绑定hello方法
//	*reply = utils.Do() //拼接字符串
//	return nil          //返回空值
//
//}

func main() {
	listen, err5 := net.Listen("tcp", "127.0.0.1:8881")
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
		var c []int
		err := json.Unmarshal(buf[:n], &c)
		if err != nil {
			fmt.Println(err)
		}
		ms := utils.Do(c)
		//var in string
		//_, _ = fmt.Scan(&in)

		con.Write([]byte(ms))
		break
	}
	//rpc.RegisterName("Slave2Service", new(Slave2Service))
	//listener, err := net.Listen("tcp", ":8881")
	//if err != nil {
	//	log.Fatal("ListenTCP erro :", err)
	//
	//}
	//
	//conn, err := listener.Accept()
	//if err != nil {
	//	log.Fatal("Accept error:", err)
	//}
	//rpc.ServeConn(conn)
}
