package main

import (
	"auto_test/utils"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"
)

//var host = []string{"43.249.9.132", "43.249.9.133"}
//var service = []string{"Slave1", "Slave2"}
//var port1 = []int{8880, 8881}

var slave = [][]string{
	//{"43.249.9.132:8880", "Slave1"},
	//{"43.249.9.133:8881", "Slave2"},
	{"127.0.0.1:8882", "Slave1"},
	{"127.0.0.1:8881", "Slave2"},
}

func main() {
	var a []int
	fmt.Println("第1轮并发数：")
	fmt.Scan(&utils.First)
	a = append(a, utils.First)
	fmt.Println("第2轮并发数")
	fmt.Scan(&utils.Second)
	a = append(a, utils.Second)
	fmt.Println("第3轮并发数")
	fmt.Scan(&utils.Third)
	a = append(a, utils.Third)
	by, _ := json.Marshal(a)
	wg := sync.WaitGroup{}
	for i, v := range slave {
		wg.Add(1)
		go func(i int, v []string) {
			conn, err5 := net.Dial("tcp", v[0])
			fmt.Println(v[0])
			if err5 != nil {
				log.Fatal(err5)
			}
			buf := make([]byte, 1024)
			fmt.Println("连接服务器成功。。。")
			for {
				//var i string
				//_, _ = fmt.Scan(&i)
				_, err5 = conn.Write(by)
				if err5 != nil {
					fmt.Println(err5)
				}
				n, err5 := conn.Read(buf)
				if err5 != nil {
					fmt.Println(err5)
				}
				log.Printf("slave[\033[32m%v\033[0m]压测结果:%v \n", v[0], string(buf[:n]))
				break
			}

			//client, err := rpc.Dial("tcp", fmt.Sprintf("%v", v[0]))
			////client, err := rpc.Dial("tcp", fmt.Sprintf("127.0.0.1:%v", port1[i]))
			//
			//if err != nil {
			//	log.Fatal("dialing : ", err)
			//}
			//var reply string
			////fmt.Println(fmt.Sprintf("%vService.%v", service[i], service[i]))
			//err = client.Call(fmt.Sprintf("%vService.%v", v[1], v[1]),
			//	fmt.Sprintf("%v", v[1]), &reply)
			//if err != nil {
			//	log.Fatal(err)
			//}
			//err1 := client.Close()
			//if err1 != nil {
			//	fmt.Println(err1)
			//}
			//fmt.Printf("slave[\033[34m%v\033[0m]压测结果：%v\n", v[0], reply)
			defer wg.Done()
		}(i, v)
	}
	wg.Wait()
	_, _ = fmt.Scanf("h")
}
