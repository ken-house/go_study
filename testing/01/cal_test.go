package cal

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := AddUpper(10)
	//if res != 55 {
	//	t.Fatalf("期望值=%v，实际值=%v\n", 55, res)
	//}
	//t.Logf("执行正确")

	// 升级使用断言
	assert.Equal(t, 55, res)
}
