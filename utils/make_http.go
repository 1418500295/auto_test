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
var Start int64
var End int64

func init() {
	cli = &http.Client{Transport: &http.Transport{
		MaxIdleConns:        10000, // Set your desired maximum number of idle connections
		MaxIdleConnsPerHost: 10000,
		IdleConnTimeout:     30 * time.Second, // Set your desired idle connection timeout
		DisableCompression:  true,             // Optional: Disable compression for testing purposes
		DisableKeepAlives:   false,            //复用连接
	}}
}

func SendPost(url string, data map[string]interface{}, token string) map[string]interface{} {
	dataBy, _ := json.Marshal(data)
	reader := bytes.NewReader(dataBy)
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("t", token)
	Start = time.Now().UnixNano() / 1e6
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	End = time.Now().UnixNano() / 1e6
	body, _ := io.ReadAll(res.Body)
	var resMap map[string]interface{}
	err = json.Unmarshal(body, &resMap)
	if err != nil {
		fmt.Println(err)
	}
	return resMap
}
