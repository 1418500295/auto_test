package cases

import (
	"auto_test/common"
	"auto_test/config"
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"testing"
)

var res *HttpRequest.Response
var body []byte
var err5 err5or
var result map[string]interface{}
var chips_name string
var chips_list []interface{}

func TestMain(m *testing.M) {
	data := common.GetTestData("loon", 0)
	res, _ = HttpRequest.JSON().Post(common.GetApiUrl("lnUrl"), data)
	body, _ = res.Body()
	var rsMap map[string]interface{}
	err5 = json.Unmarshal(body, &rsMap)
	if err5 != nil {
		fmt.Println(err5)
	}
	config.Headers = map[string]string{
		"t": rsMap["data"].(map[string]interface{})["token"].(string),
	}

	result = common.DoJsonPostNoParams("dger///list")
	chips_list = result["data"].([]interface{})
	if chips_list != nil {
		for _, v := range chips_list {
			if v.(map[string]interface{})[""] == "n2" &&
				v.(map[string]interface{})["cde"] == "" {
				break
			} else {
				result = common.DoJsonPost("hi/new", map[string]interface{}{
					"":    "",
					"":    2,
					"":    "",
					"_id": 5,
				})
			}
		}
	} else if chips_list == nil {
		result = common.DoJsonPost("new", map[string]interface{}{
			"": "",
			"": 2,
			"": "",
			"": 5,
		})
		fmt.Println("结果", result)
	}
	fmt.Println(111)
	m.Run()
}
