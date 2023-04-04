package cases

import (
	"auto_test/common"
	"auto_test/config"
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"testing"
)

func TestMain(m *testing.M) {
	data := common.GetTestData("login.json", 0)
	res, _ := HttpRequest.JSON().Post(common.GetApiUrl("loginUri"), data)
	body, _ := res.Body()
	var rsMap map[string]interface{}
	err := json.Unmarshal(body, &rsMap)
	if err != nil {
		fmt.Println(err)
	}
	config.Headers = map[string]string{
		"t": rsMap["data"].(map[string]interface{})["token"].(string),
	}
	m.Run()
}
