package cases

import (
	"auto_test/common"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestStartWorkList(t *testing.T) {
	log.Println("")
	data := common.GetTestData(".json", 0)
	res := common.DoJsonPost(common.GetApiUrl(""), data)
	assert.Equal(t, "", res["msg"])
	assert.Equal(t, , res["code"])

}
