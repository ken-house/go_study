package main

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

// 必须以Benchmark为前缀；
// b.ResetTimer是重置计时器，调用时表示重新开始计时，可以忽略测试函数中的一些准备工作
// b.N是基准测试框架提供的，表示循环的次数，因为需要反复调用测试的代码，才可以评估性能
func BenchmarkSliceAppend(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceAppend(10000)
	}
}

func BenchmarkMapAppend(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapAppend(10000)
	}
}

func TestAddUpper(t *testing.T) {
	res := AddUpper(10)
	//if res != 55 {
	//	t.Fatalf("期望值=%v，实际值=%v\n", 55, res)
	//}
	//t.Logf("执行正确")

	// 升级使用断言
	assert.Equal(t, 55, res)
}
