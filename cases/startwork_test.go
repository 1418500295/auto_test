package cases

import (
	"auto_test/common"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestStartWorkList(t *testing.T) {
	log.Println("开工列表")
	data := common.GetTestData("startwork.json", 0)
	res := common.DoJsonPost(common.GetApiUrl("startWorkListUri"), data)
	assert.Equal(t, "请求成功", res["msg"])
	assert.Equal(t, 10000, res["code"])

}
