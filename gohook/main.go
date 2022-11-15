package main

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
)

func originFunc() {
	str := "Hi, origin func"
	fmt.Println(str)
}

func monkeyFunc() {
	str := "Hi, monkey func"
	fmt.Println(str)
	trampolineFunc() // 实际调用原函数
}

// 方法体不能为空
func trampolineFunc() {
	a := 1
	fmt.Println(a)
}

func main() {
	originFunc()
	fmt.Println("-------")
	gohook.Hook(originFunc, monkeyFunc, trampolineFunc)
	originFunc()
}
