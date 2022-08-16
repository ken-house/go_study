package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 实现开启指定N个goroutine进行处理业务
func main() {
	slice := make([]int, 0, 100)
	for i := 1; i <= 100; i++ {
		slice = append(slice, i)
	}

	sChan := make(chan int, 10)

	go func() {
		for _, v := range slice {
			sChan <- v
		}
		close(sChan)
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for data := range sChan {
				fmt.Printf("goroutine num：%v，data：%v\n", i, data)
				time.Sleep(time.Second)
			}
		}(i)
	}
	wg.Wait()

}
