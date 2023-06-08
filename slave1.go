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
//func (s *Slave1Service) Slave1(request string, reply *string) error {
//	*reply = utils.Do() //将slave压测数据返回给mater
//	return nil          //返回空值
//
//}

func main() {

	lister, err := net.Listen("tcp", "127.0.0.1:8881")
	if err != nil {
		fmt.Println("listen failed error: ", err)
	}
	fmt.Printf("server start on : \033[34m%v\033[0m\n", lister.Addr().String())
	buf := make([]byte, 1024)
	for {
		con, err1 := lister.Accept()
		if err1 != nil {
			fmt.Println("accept fail error: ", err1)
		}
		n, err2 := con.Read(buf)
		if err2 != nil {
			fmt.Println("read from master error: ", err2)
		}
		var concurrencySlice []int
		err3 := json.Unmarshal(buf[:n], &concurrencySlice)
		if err3 != nil {
			fmt.Println("data covert error: ", err3)
		}
		ms := utils.ExecScript(concurrencySlice)
		//var in string
		//_, _ = fmt.Scan(&in)
		_, _ = con.Write([]byte(ms))
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
