package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var testMap = make(map[int]int, 0)
var ch = make(chan int, 1)

func test(n int) {
	defer wg.Done()
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	ch <- 1
	testMap[n] = res
	<-ch
}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()

	// 关闭ch
	close(ch)

	for i, v := range testMap {
		fmt.Printf("testMap[%v]=%v\n", i, v)
	}
}
