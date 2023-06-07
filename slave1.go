package main

import (
	"auto_test/utils"
	"encoding/json"
	"fmt"
	"net"
)

//type Slave1Service struct { //空结构体
//
//}
//
//func (s *Slave1Service) Slave1(request string, reply *string) error { //为hello结构体绑定hello方法
//	*reply = utils.Do() //拼接字符串
//	return nil          //返回空值
//
//}

func main() {

	listen, err5 := net.Listen("tcp", "127.0.0.1:8882")
	if err5 != nil {
		fmt.Println(err5)
	}
	fmt.Printf("服务已启动于: \033[34m%v\033[0m\n", listen.Addr().String())
	buf := make([]byte, 1024)
	for {
		con, err5 := listen.Accept()
		if err5 != nil {
			fmt.Println(err5)
		}
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
	//rpc.RegisterName("Slave1Service", new(Slave1Service))
	//listener, err := net.Listen("tcp", ":8880")
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
	_, _ = fmt.Scanf("h")

}
