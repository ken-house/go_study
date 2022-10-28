package main

import (
	"fmt"
	"sync"
)

var numMap = make(map[int]int, 0)
var lock sync.Mutex

// 计算一个数的阶乘并存入map中
func test(num int) {
	lock.Lock()
	res := 1
	for n := 1; n <= num; n++ {
		res *= n
	}
	numMap[num] = res
	lock.Unlock()
}

// 执行go run -race main.go可以看到有资源竞争
func main() {
	// 开启20个协程
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	// 读取map里的数据
	lock.Lock()
	for i, v := range numMap {
		fmt.Printf("数字为：%v，阶乘为：%v\n", i, v)
	}
	lock.Unlock()
}
