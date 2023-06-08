package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var cli *http.Client
var Start time.Time
var End time.Duration

func init() {
	cli = &http.Client{Transport: &http.Transport{
		MaxIdleConns:        10000, // Set your desired maximum number of idle connections
		MaxIdleConnsPerHost: 10000,
		IdleConnTimeout:     30 * time.Second, // Set your desired idle connection timeout
		DisableCompression:  true,             // Optional: Disable compression for testing purposes
		DisableKeepAlives:   false,            //复用连接
	}}
}
func reqDataConvReader(data map[string]interface{}) *bytes.Reader {
	byData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("请求body转reader错误：", err)
	}
	var reader = bytes.NewReader(byData)
	return reader
}
func SendPost(url string, data map[string]interface{}, token string) map[string]interface{} {
	req, err := http.NewRequest("POST", url, reqDataConvReader(data))
	if err != nil {
		fmt.Println("get req error: ", err)
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("t", token)
	Start = time.Now()
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println("send request error: ", err)
	}
	body, _ := io.ReadAll(res.Body)
	End = time.Since(Start)
	var resMap map[string]interface{}
	err = json.Unmarshal(body, &resMap)
	if err != nil {
		fmt.Println("convert response error: ", err)
	}
	return resMap
}
