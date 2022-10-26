package main

import (
	"fmt"
	"sync"
)

var testMap = make(map[int]int, 0)
var wg sync.WaitGroup
var mutex sync.Mutex

func test(n int) {
	defer wg.Done()
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	mutex.Lock()
	defer mutex.Unlock()
	testMap[n] = res
}

func main() {
	// 启动100个协程
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go test(i)
	}

	// 主线程等待10秒
	wg.Wait()

	for i, v := range testMap {
		fmt.Printf("testMap[%v]=%v\n", i, v)
	}
}
