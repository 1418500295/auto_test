package main

import (
	"fmt"
	"log"
	"net/rpc"
	"sync"
)

//var host = []string{"43.249.9.132", "43.249.9.133"}
//var service = []string{"Slave1", "Slave2"}
//var port1 = []int{8880, 8881}

var slave = [][]string{
	//{"43.249.9.132:8880", "Slave1"},
	//{"43.249.9.133:8881", "Slave2"},
	{"127.0.0.1:8880", "Slave1"},
	{"127.0.0.1:8881", "Slave2"},
}

func main() {
	wg := sync.WaitGroup{}
	for i, v := range slave {
		wg.Add(1)
		go func(i int, v []string) {
			client, err := rpc.Dial("tcp", fmt.Sprintf("%v", v[0]))
			//client, err := rpc.Dial("tcp", fmt.Sprintf("127.0.0.1:%v", port1[i]))

			if err != nil {
				log.Fatal("dialing : ", err)
			}
			var reply string
			//fmt.Println(fmt.Sprintf("%vService.%v", service[i], service[i]))
			err = client.Call(fmt.Sprintf("%vService.%v", v[1], v[1]),
				fmt.Sprintf("%v", v[1]), &reply)
			if err != nil {
				log.Fatal(err)
			}
			err1 := client.Close()
			if err1 != nil {
				fmt.Println(err1)
			}
			fmt.Printf("slave[\033[34m%v\033[0m]压测结果：%v\n", v[0], reply)
			defer wg.Done()
		}(i, v)
	}
	wg.Wait()
}
