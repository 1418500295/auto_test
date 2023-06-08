package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/marusama/cyclicbarrier"
	"github.com/panjf2000/ants"
	"log"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const num = 5

var et int64
var st int64

var reqUrl = "h"
var resTime int64
var chanResTime chan int64
var resTimeList []int
var rTimeChan = make(chan int, 1000000) //创建响应数据收集缓冲区
var sucNum int64
var failNum int64
var useTime int64
var avgSize int64
var totalConcurrency int
var lock = sync.Mutex{} //定义互斥锁

func avgResTime(timeList []int) float64 {
	sum := 0
	if len(timeList) == 0 {
		fmt.Println("列表数据为空")
	} else {
		for _, v := range timeList {
			sum += v
		}
	}
	return float64(sum) / float64(len(timeList))
}
func qps() float64 {
	return float64(totalConcurrency) / (avgResTime(resTimeList) / 1000)
}
func maxRespTime(timeList []int) float64 {
	max := timeList[0]
	for _, index := range timeList {
		if index > max {
			max = index
		}
	}
	return float64(max) / float64(1000)
}
func minRespTime(timeList []int) float64 {
	min := timeList[0]
	for _, index := range timeList {
		if index < min {
			min = index
		}
	}
	return float64(min) / float64(1000)
}
func fiftyRespTime(timeList []int) float64 {
	sort.Ints(timeList)
	resSize := 0.5
	return float64(timeList[int(float64(len(timeList))*resSize)-1]) / float64(1000)
}
func ninetyRespTime(timeList []int) float64 {
	sort.Ints(timeList)
	resSize := 0.9
	return float64(timeList[int(float64(len(timeList))*resSize)-1]) / float64(1000)
}

var reqData = map[string]interface{}{
	"page": 1,
	"size": 20,
}
var headers = map[string]string{
	"t": "",
}

//func getData() *bytes.Reader {
//	dataBy, _ := json.Marshal(reqData)
//	reader := bytes.NewReader(dataBy)
//	return reader
//}

func execute(i int, times int64) {
	defer func() {
		if err5 := recover(); err5 != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			fmt.Printf("safe go: %v\n%v\n", err5, *(*string)(unsafe.Pointer(&buf)))
		}
	}()
	st = time.Now().UnixMilli()
	for true {
		res := SendPost(reqUrl, reqData, headers["t"])
		Times := int64(End)
		fmt.Println("请求耗时：", End)
		rTimeChan <- int(Times / 1e6) //并发时安全写入数据
		log.Printf("第%d个协程返回：%v\n", i, res["code"])
		if res["msg"].(string) == `请求成功` && res["code"].(float64) == 10000 {
			atomic.AddInt64(&sucNum, 1) //原子操作，保证并发时数据递增的安全性
		} else {
			atomic.AddInt64(&failNum, 1)
		}
		et = time.Now().UnixMilli()
		if et-st > times*1000 {
			break
		}
		continue
	}
}

func run(num int, executeTimes int64) {
	defer func() { //进行详细的异常捕获
		if err5 := recover(); err5 != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			fmt.Printf("safe go: %v\n%v\n", err5, *(*string)(unsafe.Pointer(&buf)))
		}
	}()
	barrier := cyclicbarrier.New(num)
	wg := sync.WaitGroup{}
	pool, _ := ants.NewPool(num)
	for i := 0; i < num; i++ {
		c := i
		wg.Add(1)
		err := pool.Submit(func() {
			err5 := barrier.Await(context.Background()) //集合点同步
			if err5 != nil {
				fmt.Println("集合点同步异常:", err5)
			}
			execute(c, executeTimes)
			defer wg.Done()
		})
		if err != nil {
			return
		}
	}
	wg.Wait() //主线程等待子线程执行完毕，收集数据
	defer func() {
		_ = pool.Release()
	}()
}

func printTable(num int) string {
	table, err5 := gotable.Create("耗时", "并发数", "成功的请求数",
		"失败的请求数", "平均响应时间", "50%响应时间", "90%响应时间", "最大响应时间", "最小响应时间", "qps")
	if err5 != nil {
		fmt.Println(err5)
	}
	var result = make(map[string]interface{})
	useTime = et - st
	err := table.AddRow([]string{fmt.Sprintf("%.1f", float64(useTime/1000)),
		strconv.Itoa(num), strconv.FormatInt(sucNum, 10),
		strconv.FormatInt(failNum, 10),
		fmt.Sprintf("%.3f", avgResTime(resTimeList)/1000),
		fmt.Sprintf("%.3f", fiftyRespTime(resTimeList)),
		fmt.Sprintf("%.3f", ninetyRespTime(resTimeList)),
		fmt.Sprintf("%.3f", maxRespTime(resTimeList)),
		fmt.Sprintf("%.3f", minRespTime(resTimeList)),
		fmt.Sprintf("%.3f", qps()),
	})
	if err != nil {
		fmt.Println(err)
	}
	result["耗时"] = fmt.Sprintf("%.1f", float64(useTime/1000))
	result["并发数"] = strconv.Itoa(num)
	result["成功请求数"] = strconv.FormatInt(sucNum, 10)
	result["失败请求数"] = strconv.FormatInt(failNum, 10)
	result["平均响应时间"] = fmt.Sprintf("%.3f", avgResTime(resTimeList)/1000)
	result["50%响应时间"] = fmt.Sprintf("%.3f", fiftyRespTime(resTimeList))
	result["90%响应时间"] = fmt.Sprintf("%.3f", ninetyRespTime(resTimeList))
	result["最大响应时间"] = fmt.Sprintf("%.3f", maxRespTime(resTimeList))
	result["最小响应时间"] = fmt.Sprintf("%.3f", minRespTime(resTimeList))
	result["QPS"] = fmt.Sprintf("%.3f", qps())
	fmt.Println(table)
	by, _ := json.Marshal(result)
	return string(by)
}

func ExecScript(concurrencySlice []int) string {
	totalConcurrency = 0
	for i, concurrency := range concurrencySlice {
		fmt.Printf("5秒后进行第 %d 轮并发,并发数为：%d\n", i+1, concurrency)
		<-time.After(5 * time.Second)
		run(concurrency, 2)
		totalConcurrency += concurrency
	}
	close(rTimeChan) //关闭通道
	for ch := range rTimeChan {
		resTimeList = append(resTimeList, ch)
	}
	return printTable(totalConcurrency)
}
