package cases

import (
	"auto_test/common"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestEntryRecord(t *testing.T) {
	log.Println("入场记录")
	data := common.GetTestData(".json", 0)
	res := common.DoJsonPost(common.GetApiUrl(""), data)
	assert.Equal(t, "请求成功", res["msg"])
	assert.Equal(t, 10000, res["code"])
}
