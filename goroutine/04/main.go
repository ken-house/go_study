package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

func main() {
	// 初始化goroutine协程池
	pool, _ := ants.NewPool(2)
	var wg sync.WaitGroup
	// 定义一个goroutine要执行的方法
	syncCalculateSum := func() {
		func() {
			defer wg.Done()
			fmt.Println("hello world")
			time.Sleep(1 * time.Second)
		}()
	}
	// 执行函数十次，但仅使用两个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		pool.Submit(syncCalculateSum)
	}
	wg.Wait()
	// 释放goroutine协程池
	pool.Release()
}
