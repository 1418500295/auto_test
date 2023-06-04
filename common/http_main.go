package common

import (
	"auto_test/config"
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"log"
)

func DoPost(url string, testData map[string]interface{}) map[string]interface{} {
	res, err5 := HttpRequest.Post(url, testData)
	if err5 != nil {
		fmt.Println(err5)
	}
	body, _ := res.Body()
	defer res.Close()
	defer log.Println("响应时间：" + fmt.Sprintf("%v", res.Time()))
	log.Printf("接口返回：%v", string(body))
	var jsonResp map[string]interface{}
	err7 := res.Json(&jsonResp)
	if err7 != nil {
		fmt.Println(err7)
	}
	//判断map中的值类型，因为map解析会将int值自动转为float64，并将值(一般为interface类型)转为string或int
	for k, v := range jsonResp {
		switch v.(type) {
		case float64:
			jsonResp[k] = int(v.(float64))
		case string:
			jsonResp[k] = v.(string)
		}

	}
	defer res.Close()
	return jsonResp
}

func DoJsonPost(url string, testData map[string]interface{}) map[string]interface{} {
	res, _ := HttpRequest.SetHeaders(config.Headers).JSON().Post(url, testData)
	body, _ := res.Body()
	defer res.Close()
	//defer log.Println("响应时间：" + fmt.Sprintf("%v", res.Time()))
	//log.Printf("接口返回：%v", string(body))
	var jsonRes map[string]interface{}
	err5 := json.Unmarshal(body, &jsonRes)
	if err5 != nil {
		fmt.Println(err5)
	}
	for k, v := range jsonRes {
		switch v.(type) {
		case float64:
			jsonRes[k] = int(v.(float64))
		case string:
			jsonRes[k] = v.(string)
		}
	}
	return jsonRes
}

func DoJsonPostNoParams(url string) map[string]interface{} {
	res, _ := HttpRequest.SetHeaders(config.Headers).JSON().Post(url)
	body, _ := res.Body()
	defer res.Close()
	//defer log.Println("响应时间：" + fmt.Sprintf("%v", res.Time()))
	//log.Printf("接口返回：%v", string(body))
	var jsonRes map[string]interface{}
	err5 := json.Unmarshal(body, &jsonRes)
	if err5 != nil {
		fmt.Println(err5)
	}
	for k, v := range jsonRes {
		switch v.(type) {
		case float64:
			jsonRes[k] = int(v.(float64))
		case string:
			jsonRes[k] = v.(string)
		}
	}
	return jsonRes
}

func DoGet(url string, testData map[string]interface{}) map[string]interface{} {
	res, err5 := HttpRequest.Get(url, testData)
	if err5 != nil {
		fmt.Println(err5)
	}
	body, _ := res.Body()
	defer res.Close()
	defer log.Println("响应时间: " + fmt.Sprintf("%v", res.Time()))
	log.Printf("接口返回：%v", string(body))
	var jsonRes map[string]interface{}
	err7 := json.Unmarshal(body, &jsonRes)
	if err7 != nil {
		fmt.Println(err7)
	}
	for k, v := range jsonRes {
		switch v.(type) {
		case float64:
			jsonRes[k] = int(v.(float64))
		case string:
			jsonRes[k] = v.(string)
		}
	}
	return jsonRes

}
func DoGetNoParams(url string) map[string]interface{} {
	res, err5 := HttpRequest.Get(url)
	if err5 != nil {
		fmt.Println(err5)
	}
	body, _ := res.Body()
	defer res.Close()
	defer log.Println("响应时间：" + fmt.Sprintf("%v", res.Time()))
	log.Printf("接口返回：%v", string(body))
	var jsonRes map[string]interface{}
	err6 := json.Unmarshal(body, &jsonRes)
	if err6 != nil {
		fmt.Println(err6)
	}
	for k, v := range jsonRes {
		switch v.(type) {
		case float64:
			jsonRes[k] = int(v.(float64))
		case string:
			jsonRes[k] = v.(string)
		}
	}
	return jsonRes
}
