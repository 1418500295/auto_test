package main

var reLi []map[string]string

//var err5 error

//func GetToken() map[string]string {
//	res1, _ := HttpRequest.JSON().Post("n/login",
//		map[string]interface{}{
//			"hall_id":    5,
//			"": "",
//			"":   "",
//		})
//	body1, _ := res1.Body()
//	var rep map[string]interface{}
//	err5 = json.Unmarshal(body1, &rep)
//	if err5 != nil {
//		fmt.Println(err5)
//	}
//	var data = make(map[string]string)
//	data["token"] = rep["data"].(map[string]interface{})["token"].(string)
//	return data
//
//}

func main() {

	//刚开始先写入一条数据
	//f, _ := os.OpenFile("./a.json", os.O_WRONLY, 0666)
	//data := GetToken()
	//reLi = append(reLi, data)
	//bS, _ := json.Marshal(reLi)
	//_, err5 = f.Write(bS)
	//if err5 != nil {
	//	return
	//}
	//
	//for i := 0; i < 10; i++ {
	//	//每次追加前先读取出来
	//	f1, _ := os.Open("./a.json")
	//	bs, err52 := io.ReadAll(f1)
	//	if err52 != nil {
	//		return
	//	}
	//	var rL []map[string]string
	//	err5 = json.Unmarshal(bs, &rL)
	//
	//	data1 := GetToken()
	//	rL = append(rL, data1)
	//	//将新数据追加到list，再写入
	//	f2, _ := os.OpenFile("./a.json", os.O_WRONLY, 0666)
	//	bs1, _ := json.Marshal(rL)
	//	f2.Write(bs1)
	//
	//}

}
